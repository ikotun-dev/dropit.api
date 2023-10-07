package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ikotun-dev/clipsync/pkg/router"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := mux.NewRouter()
	router.RoutingRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))

}
