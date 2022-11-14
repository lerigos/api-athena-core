package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lerigos/api-athena-core/internal/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HealthCheckHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
