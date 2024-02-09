package models

import (
	"context"
	"fmt"

	"github.com/swarnikaraj/go-mongo-book-crud/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
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
	var Books []Book
	collection.Find(context.Background(),&Books)
	return Books
}

func ( b *Book) BookUpdater(updateFields map[string]interface{}, Id string) (*Book , error){

	id,_:=primitive.ObjectIDFromHex(Id)
    filter:=bson.M{"_id":id}
	update:=bson.M{"$set":updateFields}
	result:=collection.FindOneAndUpdate(context.Background(),filter,update)
	if result.Err() != nil {
        return nil, result.Err()
    }
    var book =&Book{} 
	err:= result.Decode(book)
	if err!=nil{
		return nil, err
	}
	return book, nil
}