package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/haroonalbar/book-store/pkg/routes"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	routes.ResgisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Listening on port ", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, r))
}
