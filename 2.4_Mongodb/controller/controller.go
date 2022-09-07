package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/dewasahu2003/mongodbAPI/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// we will define two set of helper method
//one for client side processing
//and second for mongoDB for CRUD operations

//mongoDB Insert 1 record

func mongo_insertMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("movie inserted in db with id:%v", inserted.InsertedID)
}

func mongo_updateMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}                         //searching criteria
	update := bson.M{"$set": bson.M{"iswatched": true}} //updating value

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The _id updated is:%+v", result.ModifiedCount)

}

func mongo_deleteMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("record id:%+v\n delete count:%+v", id, deleteResult)
}

func mongo_deleteAllMovie() {
	filter := bson.D{{}}

	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("all records deleted:%+v", deleteResult.DeletedCount)
}
