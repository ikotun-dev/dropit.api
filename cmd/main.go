package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ikotun-dev/clipsync/pkg/middleware"
	"github.com/ikotun-dev/clipsync/pkg/router"
	"github.com/ikotun-dev/clipsync/pkg/socket"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.LogRequest)
	router.RoutingRoutes(r)
	r.HandleFunc("/ws", socket.WebSocketHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))

}
