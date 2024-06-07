// routes required controller
// app.go requireed gormmodels required config

package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sahana/go-bookstore/pkg/config"
)

var db *gorm.DB

// model provide struture help to store in db
type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

//whenerver we are working db we will edit init function

func init() {
	config.Connect()    //connect db
	db = config.GetDB() //whterv is get from calling function we get from app.go
	db.AutoMigrate(&Book{})
}

//user interact with routes it send contro, to controller it has logic and controller has to perfon operation with db operation of db has to reside in model it has to have different function for different controller
//we need to have createbook function here and get book,delete book
//either create func diretly or in main.go FILE
//create controller then create model func (talk to db)and then create routes
//frist model func then corespornding controller then routes

// craete func that need to talk to db
// db will talk to database
func (b *Book) CreateBook() *Book {
	db.NewRecord(b) //exist inside gorm
	db.Create(&b)
	return b

}

func GetAllBooks() []Book {
	var Books []Book //list book in db and return
	db.Find(&Books)
	return Books
}

// id is to be taken
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook) //where and find that book
	return &getBook, db                       //founded book and db variable
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
