package main

import (
	"log"
	"net/http"

	"github.com/edsistant/backend/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/", api.Index).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
