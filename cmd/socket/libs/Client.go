package libs

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	MaxMessageSize = 1024 * 1024
)

type Client struct {
	WS   *websocket.Conn
	Send chan []byte
}

/*
 * @desc Function for get real time message from websocket
 *
 */
func (c *Client) ReadPump() {
	defer func() {
		H.UnRegister <- c
		c.WS.Close()
	}()

	c.WS.SetReadLimit(MaxMessageSize)
	c.WS.SetReadDeadline(time.Now().Add(pongWait))
	c.WS.SetPongHandler(func(string) error {
		c.WS.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.WS.ReadMessage()
		if err != nil {
			break
		}
		H.Broadcast <- string(message)
	}
}

/*
 * @desc Function for display data to websocket
 *
 */
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.WS.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *Client) write(mt int, message []byte) error {
	c.WS.SetWriteDeadline(time.Now().Add(writeWait))
	return c.WS.WriteMessage(mt, message)
}
