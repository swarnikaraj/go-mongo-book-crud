package models

import (
	"context"
	"fmt"
	"log"

	"github.com/swarnikaraj/go-mongo-book-crud/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

type Book struct{ 
 Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
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


func (b *Book) BooksGetter() []primitive.M{
	var Books []primitive.M
	cursor,err:=collection.Find(context.Background(),bson.D{{}})

	  if err!=nil{
		log.Fatal(err)
	  }
		for cursor.Next(context.TODO()){
			var book bson.M
           if err:=cursor.Decode(&book); err!=nil{
			log.Fatal(err)
		   }
		   Books=append(Books, book)
		}
	defer cursor.Close(context.Background())
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


func (b *Book) DeleteOne(Id string) (int64){
   id,_:=primitive.ObjectIDFromHex(Id)
   filter:=bson.M{"_id":id}
   result,err:=collection.DeleteOne(context.Background(),filter)
if err!=nil{
log.Fatal(err)
}
   
   return result.DeletedCount
}

func (b *Book) DeleteAllRecords() int64{
	result, err:=collection.DeleteMany(context.Background(),bson.D{{}},nil)
	if err!=nil{
		log.Fatal(err)
	}
	return result.DeletedCount
}


