package controllers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	rooms      = make(map[string]map[*websocket.Conn]struct{})
	roomsMtx   sync.Mutex
	socketsMtx sync.Mutex
)

func reader(conn *websocket.Conn, key string) {
	defer func() {
		// Remove the client from the room when they disconnect
		roomsMtx.Lock()
		delete(rooms[key], conn)
		roomsMtx.Unlock()
		conn.Close()
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Received:", string(p))

		// Broadcast the message to all connected clients in the same room
		roomsMtx.Lock()
		for client := range rooms[key] {
			if err := client.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}
		roomsMtx.Unlock()
	}
}

func SocketEndpoint(w http.ResponseWriter, r *http.Request) {
	// Get the session key from the query string
	key := r.URL.Query().Get("session_key")

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.Write([]byte("Error occurred"))
		return
	}

	// Add the socket to the room based on the session key
	roomsMtx.Lock()
	if _, ok := rooms[key]; !ok {
		rooms[key] = make(map[*websocket.Conn]struct{})
	}
	rooms[key][ws] = struct{}{}
	roomsMtx.Unlock()

	//returns socket number and user
	log.Println("Client connected to socket room  --> : ", key)

	reader(ws, key)
}
