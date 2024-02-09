package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swarnikaraj/go-mongo-book-crud/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	bookRoutes:=r.PathPrefix("/books").Subrouter()
	routes.BookRouter(bookRoutes)

	r.HandleFunc("/", func(w http.ResponseWriter,r *http.Request){
     fmt.Fprintf(w,"This is Test Route")
	})
	// r.Handle("/books",r)
    fmt.Print("Server is Running on port 5000")
	err:= http.ListenAndServe(":5000",r)

	if err!=nil{
		log.Fatal(err)
	}

	
}