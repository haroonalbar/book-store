package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haroonalbar/book-store/pkg/routes"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()

	routes.ResgisterBookStoreRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
