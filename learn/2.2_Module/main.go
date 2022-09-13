package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//go mod tidy := make the mod file direct
	//go mod verify := to verify the go.sum and verify modules that we imported
	//go list -m all := to find the packages that our project is
	// go list -m -versions github.com/gorilla/mux:= to directly get the version tree
	// go mod graph := to get dependency tree
	// go mod vendor := to override things in the package that you are using and to run with this
	//use -> go run -mod=vendor main.go

	fmt.Println("MODULES")
	r := mux.NewRouter()
	r.HandleFunc("/", Server).Methods("GET")

	//http.ListenAndServe(":1000",r) -- can give error

	log.Fatal(http.ListenAndServe(":1000", r))

}
func Server(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>CHANGE THE ORDER </h1>"))
}
