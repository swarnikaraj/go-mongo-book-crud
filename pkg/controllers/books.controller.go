package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swarnikaraj/go-mongo-book-crud/pkg/models"
	"github.com/swarnikaraj/go-mongo-book-crud/pkg/utils"
)
var BookModel models.Book
func CreateBook(w http.ResponseWriter, r *http.Request){

	var userInstance =&models.Book{}

	utils.BodyParser(r, userInstance)
	createtionResponse:= userInstance.BookCreator()
	res, _:=json.Marshal(createtionResponse)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}


func GetBooks(w http.ResponseWriter, r *http.Request){
	books:=BookModel.BooksGetter()
	fmt.Print("books getter", books)
	res, _:=json.Marshal(books)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var book =models.Book{}
	vars:=mux.Vars(r)
    bookId:=vars["id"]
	
    var fieldUpdate map[string]interface{} 
   utils.BodyParser(r,&fieldUpdate)
	res, err:=book.BookUpdater(fieldUpdate,bookId)
	if err!=nil{
		w.Header().Set("Content-Type","application/json")
	  w.WriteHeader(http.StatusBadRequest)
     return
	}
    updatedBookJSON, _ := json.Marshal(res)
	
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusAccepted)
    w.Write(updatedBookJSON)

}


func DeleteOneBook(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	Id:=vars["id"]
	delRes:=BookModel.DeleteOne(Id)
    res,_:=json.Marshal(delRes)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}

func DeleteAll (w http.ResponseWriter,r *http.Request){
	delRes:=BookModel.DeleteAllRecords()
	res,_:=json.Marshal(delRes)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

