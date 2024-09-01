package models

import (
	"log"

	"github.com/haroonalbar/book-store/pkg/config"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// !look
var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env %v", err)
	}
	config.Connect()
	db = config.GetDB()
	// !look
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// !look
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	// !look
	db.Find(&Books)
	return Books
}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var getBook Book
	// !look
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	// !look
	db.Where("ID=?", id).Delete(&book)
	return book
}
