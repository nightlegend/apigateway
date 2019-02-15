package socketio

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WShandler web socket handler
// accpect client websocket connection
func WShandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	conns = append(conns, conn)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(messageType, message)
	}
}
