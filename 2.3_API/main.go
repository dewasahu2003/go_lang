package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
func (c Courses) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("API")
}

//controller

//home route

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>change the order</h1>"))
}

func allcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
