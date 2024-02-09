package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://swarnikarajsingh:n3oWFCD5NO0rTTnV@cluster0.twmi8ps.mongodb.net/?retryWrites=true&w=majority"
// var collection *mongo.Collection
var db *mongo.Client
func ClientcreateDbConnection(){
clientoption := options.Client().ApplyURI(connectionString)

	client, err:=mongo.Connect(context.TODO(),clientoption)


	if err!=nil{
		log.Fatal(err)
	}
	db=client
}


func GetDbConnection( ) *mongo.Client{

	return db
}