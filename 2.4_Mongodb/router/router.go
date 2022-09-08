package router

import (
	"github.com/dewasahu2003/mongodbAPI/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	controller.Init()
	router := mux.NewRouter()

	router.HandleFunc("/api/movie", controller.Add_movie).Methods("POST")             //add
	router.HandleFunc("/api/movies", controller.Get_allMovies).Methods("GET")         //get all
	router.HandleFunc("/api/movie/{id}", controller.Update_movie).Methods("PUT")      //update
	router.HandleFunc("/api/movies/{id}", controller.Delete_movie).Methods("DELETE")  //delete
	router.HandleFunc("/api/deleteall", controller.Delete_ALLmovie).Methods("DELETE") //delete all

	return router
}
