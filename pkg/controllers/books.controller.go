package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/swarnikaraj/go-mongo-book-crud/pkg/models"
	"github.com/swarnikaraj/go-mongo-book-crud/pkg/utils"
)
var book models.Book
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
	books:=book.BooksGetter()
	res, _:=json.Marshal(books)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}