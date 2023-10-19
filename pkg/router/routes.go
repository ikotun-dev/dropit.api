package router

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/ikotun-dev/clipsync/pkg/controllers"
)
allowedOrigins := handlers
var RoutingRoutes = func(router *mux.Router) {
	router.HandleFunc("/session/new", controllers.CreateSession).Methods("POST")
	router.HandleFunc("/session/ws", controllers.JoinSession).Methods("POST")
}
