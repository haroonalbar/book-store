package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haroonalbar/book-store/pkg/controllers"
)

var ResgisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/book/", controllers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", controllers.GetBookByID).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods(http.MethodDelete)
}
