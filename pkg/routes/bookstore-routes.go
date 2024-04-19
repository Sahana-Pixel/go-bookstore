// 1
package routes

import (
	"github.com/Sahana-Pixel/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

// routes for the book store
// routes will get to controller to meet my logic
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST") //handlers is somebody come to book ,route i want my controller.creatbook to handle
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
