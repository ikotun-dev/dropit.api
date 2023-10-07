package router

import (
	"github.com/gorilla/mux"
	"github.com/ikotun-dev/clipsync/pkg/controllers"
)

var RoutingRoutes = func(router *mux.Router) {
	router.HandleFunc("/session/new", controllers.CreateSession).Methods("POST")
}
