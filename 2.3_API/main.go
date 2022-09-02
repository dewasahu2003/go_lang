package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Courses struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

//fake database
var courses []Courses

//middleware,helper
func (c *Courses) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("API")
}

//controller

//home route
// w:= POST type of method
//r:= GET type of method
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>change the order</h1>"))
}

func getallcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func get_oneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🎊-Getting one course")
	w.Header().Set("Content-Type", "application/json")

	//getting what is passed in REQUEST
	params := mux.Vars(r)
	fmt.Println("params: ", params)

	//querying the db
	for _, course := range courses {
		if course.CourseId == params["id"] {
			//give a RESPONSE
			json.NewEncoder(w).Encode(course) //giving |course| as response
			return
		} else {
			json.NewEncoder(w).Encode("no course found 💃")
			return
		}

	}

}
