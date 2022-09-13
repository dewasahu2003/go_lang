package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dewasahu2003/mongodbAPI/router"
)

func main() {
	fmt.Println("Mongodb🌿")
	r := router.Router()

	fmt.Println("THE SERVER HAS STARTED🔥🚀🚨☀️☄️⚡")
	log.Fatal(http.ListenAndServe(":1000", r))

}
