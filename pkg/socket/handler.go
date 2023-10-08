package socket

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/ikotun-dev/clipsync/pkg/controllers"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer conn.Close()

	controllers.CreateSession(w, r, conn)
}
