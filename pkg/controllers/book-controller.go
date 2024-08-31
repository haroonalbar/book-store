package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/haroonalbar/book-store/pkg/models"
	"github.com/haroonalbar/book-store/pkg/utils"
)

var Book models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Error while parsing id", http.StatusBadRequest)
		return
	}
	book, _ := models.GetBookByID(id)
	if err != nil {
		http.Error(w, "Error getting book", http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	// from utils
	utils.ParseBody(r, book)
	// from models
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	id, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		http.Error(w, "Error while parsing id", http.StatusBadRequest)
		return
	}
	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	// read from body and put it on book
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookID := vars["id"]
	id, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		http.Error(w, "Error while parsing id", http.StatusBadRequest)
		return
	}
	book, db := models.GetBookByID(id)

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}
	// book.UpdatedAt = time.Now().UTC()

	db.Save(&book)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
