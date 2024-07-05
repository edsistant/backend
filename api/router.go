package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", index)

	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Edsistant API v1"))
}
