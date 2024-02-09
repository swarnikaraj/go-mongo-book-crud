package routes

import (
	"github.com/gorilla/mux"
	"github.com/swarnikaraj/go-mongo-book-crud/pkg/controllers"
)

var BookRouter = func(router *mux.Router){
	router.HandleFunc("/getAll",controllers.GetBooks).Methods("GET")
	router.HandleFunc("/create", controllers.CreateBook).Methods("POST")
    router.HandleFunc("/update/{id}",controllers.UpdateBook).Methods("PATCH")
	router.HandleFunc("/deleteOne/{id}", controllers.DeleteOneBook).Methods("DELETE")
}