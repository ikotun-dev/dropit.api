package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/ikotun-dev/clipsync/pkg/controllers"
)

allowedOrigins := handlers.AllowedOrigins([]string{"*"})
allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})
allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

corsHandler := handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)(r)
var RoutingRoutes = func(router *mux.Router) {
	router.HandleFunc("/session/new", controllers.CreateSession).Methods("POST")
	router.HandleFunc("/session/ws", controllers.JoinSession).Methods("POST")
}
