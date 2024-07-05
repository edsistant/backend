package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/edsistant/backend/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	mount(r, "/api", api.Router())
	log.Fatal(http.ListenAndServe(":8080", r))
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
