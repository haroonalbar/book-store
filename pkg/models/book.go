package models

import (
	"log"

	"github.com/haroonalbar/book-store/pkg/config"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

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
	//  This function will create tables, missing columns, and missing indexes, and will not delete/change any existing columns and their types
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// checks if the record is new or not
	db.NewRecord(b)
	// add the record
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
