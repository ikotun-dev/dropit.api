package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ikotun-dev/clipsync/pkg/controllers"
	"github.com/ikotun-dev/clipsync/pkg/middleware"
	"github.com/ikotun-dev/clipsync/pkg/router"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Create a new router and apply middleware.
	r := mux.NewRouter()
	r.Use(middleware.LogRequest)

	// Define your allowed origins, headers, and methods.
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	// Apply CORS middleware to the router.
	corsHandler := handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)

	// Define your routes and apply the CORS middleware.
	router.RoutingRoutes(r)

	// Handle WebSocket requests.
	r.HandleFunc("/ws", controllers.SocketEndpoint)

	// Apply CORS to the entire router.
	http.Handle("/", corsHandler(r))

	// Start the server.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
