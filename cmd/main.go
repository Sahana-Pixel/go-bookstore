// will create server or will tell where router reside
package main

import (
	"log" //logout if there is an error
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //for mysql

	//import routes ,routes inside pkg
	"github.com/sahana/go-bookstore/pkg/routes"
)

func main() {
	//intailase the router here
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) //roues we have registerbookstore and we are pssing r
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) //craete server which is passed with localhost
	//localhost is port or address
	//its insdie http package
	//if there is error ther is log8

}
