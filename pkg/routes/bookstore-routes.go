package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haroonalbar/book-store/pkg/controllers"
)

var ResgisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/book/", controllers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/book/{bookID}", controllers.GetBookByID).Methods(http.MethodGet)
	router.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods(http.MethodDelete)
}
