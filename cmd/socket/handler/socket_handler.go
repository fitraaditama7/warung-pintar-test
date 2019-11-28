package handler

import (
	"net/http"
	"warung-pintar-test/cmd/socket/libs"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  libs.MaxMessageSize,
	WriteBufferSize: libs.MaxMessageSize,
}

/*
 * GET : '/ws'
 *
 * @desc Websocket for realtime message
 *
 */
func ServerWs(c *gin.Context) {
	ServerWS(c.Writer, c.Request)
}

func ServerWS(w http.ResponseWriter, r *http.Request) {

	Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("Unable to create websocket connection got error: %v", err)
		return
	}

	c := &libs.Client{
		Send: make(chan []byte, libs.MaxMessageSize),
		WS:   ws,
	}

	libs.H.Register <- c

	go c.WritePump()
	c.ReadPump()
}
