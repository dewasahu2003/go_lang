package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	//Encoder-> turing our object into json  --via (w) 💯
	//Decoder-> turing our json into object --via (r) 💯
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

func add_oneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🎊-adding one course")
	w.Header().Set("Content-Type", "application/json")

	// what if := body of req is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("The response is invalid --empty response")
	}

	//what if := body of req is {}
	var course *Courses
	_ = json.NewDecoder(r.Body).Decode(course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside JSON")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, *course)
	json.NewEncoder(w).Encode(course)
	return
}
func update_OneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🎊-updating one course")
	w.Header().Set("Content-Type", "application/json")

	//get the course needed to be updated -> get id

	params := mux.Vars(r)

	//find the id in db,delete it,add new course with same id

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) //delete

			var new_course Courses //make new object
			_ = json.NewDecoder(r.Body).Decode(&new_course)

			new_course.CourseId = params["id"] //keep the id same

			courses = append(courses, new_course) //add to db

			json.NewEncoder(w).Encode(&new_course) //POST it
			return
		}
	}
}

func delete_OneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🎊-deleting one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break //put break it will stop the looping then and there itself
		}
	}
}
