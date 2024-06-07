// 1
package routes

import (
	//configfile needed to connect db
	//"github.com/sahana/go-bookstore/pkg/config"
	"github.com/gorilla/mux"
	"github.com/sahana/go-bookstore/pkg/controllers"
)

// routes for the book store
// routes will get to controller to meet my logic
// passed r from main.go to router of type mux
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	//we handle function using controller function
	//handlers is somebody come to book ,route i want my controller.creatbook to handle
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
