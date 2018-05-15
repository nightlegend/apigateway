package socketio

import (
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/googollee/go-socket.io"
)

// RunServer is start a socket server.
func RunServer() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Info("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			log.Info("emit:", so.Emit("chat message", msg))
			log.Info(msg)
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			log.Info("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Info("error:", err)
	})

	http.Handle("/socket.io/", server)
	log.Info("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
