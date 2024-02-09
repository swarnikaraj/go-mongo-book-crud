package models

import (
	"context"
	"fmt"

	"github.com/swarnikaraj/go-mongo-book-crud/pkg/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

type Book struct{ 
 Id primitive.ObjectID `json:"_id omitempty" bson:"_id omitempty"`
 Name string `json:"name"`
 Author string `json:"author"`
}



func init() {
	 config.ClientcreateDbConnection()
	 client:=config.GetDbConnection()
     collection=client.Database("gobookcrud").Collection("book")

	 fmt.Print("Collection reference is ready")
}


func (b *Book) BookCreator() *Book{
	collection.InsertOne(context.Background(),&b)
	
	return b
}


func (b *Book) BooksGetter() []Book{
	var books []Book
	collection.Find(context.Background(),&books)
	return books
}