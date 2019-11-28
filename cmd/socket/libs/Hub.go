package libs

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan string
	Register   chan *Client
	UnRegister chan *Client
	Content    string
}

var H = Hub{
	Clients:    make(map[*Client]bool),
	Broadcast:  make(chan string),
	Register:   make(chan *Client),
	UnRegister: make(chan *Client),
	Content:    "",
}

/*
 * @desc Function pool for websocket
 *
 */
func (h *Hub) Run() {
	for {
		select {
		case c := <-h.Register:
			h.Clients[c] = true
			c.Send <- []byte(h.Content)
			break

		case c := <-h.UnRegister:
			_, ok := h.Clients[c]
			if ok {
				delete(h.Clients, c)
				close(c.Send)
			}
			break

		case m := <-h.Broadcast:
			h.Content = m
			h.broadcastMessage()
			break
		}
	}
}

/*
 * @desc Function blast message to websocket
 *
 */
func (h *Hub) broadcastMessage() {
	for c := range h.Clients {
		select {
		case c.Send <- []byte(h.Content):
			break
		default:
			close(c.Send)
			delete(h.Clients, c)
		}
	}
}
