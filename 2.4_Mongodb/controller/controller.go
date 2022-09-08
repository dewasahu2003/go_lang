package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dewasahu2003/mongodbAPI/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://dewa7sahu:dewa7sahu@cluster1.nr2jhgi.mongodb.net/?retryWrites=true&w=majority"
const dbname = "netflix"
const colName = "watchlist"

//MOST IMPORTANT
var collection *mongo.Collection

//connect with mongoDB
func Init() {
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

func mongo_getAllMovie() []primitive.M {
	// {{}} := means we dont have any filter -> means we select every thing
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	// cursor:= is like iterator

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M

		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

//methods for routing --actual 💃

func Get_allMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allmovies := mongo_getAllMovie()
	json.NewEncoder(w).Encode(allmovies)

}

func Add_movie(w http.ResponseWriter, r *http.Request) {
	//taking request from the client (front-end)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix //giving it a shape

	_ = json.NewDecoder(r.Body).Decode(&movie) //putting the data in a shape
	mongo_insertMovie(movie)                   //adding the movie to mongoDB
	json.NewEncoder(w).Encode(movie)

}

func Update_movie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r) //taking the id of movie that we need to change

	mongo_updateMovie(params["id"]) //change it in db
	json.NewEncoder(w).Encode(params["id"])

}
func Delete_movie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	mongo_deleteMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func Delete_ALLmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	mongo_deleteAllMovie()
	json.NewEncoder(w).Encode("Delete-all-Movies")
}
