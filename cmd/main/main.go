package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tokha04/go-bookstore-mysql-gorilla/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
