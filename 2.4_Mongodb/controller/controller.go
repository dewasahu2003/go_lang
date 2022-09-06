package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://dewa7sahu:dewa7sahu@cluster0.gnvuxfo.mongodb.net/?retryWrites=true&w=majority"
const dbname = "netflix"
const colName = "watchlist"

//MOST IMPORTANT
var collection *mongo.Collection

//connect with mongoDB
func init() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongoDb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	//context.Background():= shall keep a connection in backgground
	///context.TODO():= shall connect only one time

	if err != nil {
		log.Panic(err)
	}
	fmt.Println("connected to MongoDB🍃 succesful")

	collection = client.Database(dbname).Collection(colName)
	//passing the collection as var

}
