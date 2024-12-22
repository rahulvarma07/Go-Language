// In this let's understand how to build our own API's

package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Initially we did not work with DataBase
// So lets assume these are model(structure) of our database..

// Models - (file)
type coureses struct {
	CourseName  string       `json:"coursename"`
	CourseID    string       `json:"course_id"`
	CourseFee   float32      `json:"course_fee"`
	CourseTutor *courseTutor `json:"course_tutor"` // Model interaction
}

// Model 2
type courseTutor struct {
	TutorName string `json:"tutor_name"`
}

// Model-end;

// Fake DataBase
var allCourses []coureses // Storing the model's inside a struct..

// Let's define some helper function
// helper function (used to valid the function)

// This function helps for checking whethre the course ID and
// name are not empty..

func (course *coureses) validIDAndName() bool {
	//return (course.CourseName != "" && course.CourseID != "") (commenting because i want to generate CourseID)
	return course.CourseName != ""
}

func main() {

}

// When Opened a local host ensures there is some HTML code
// Contrller's
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Building An API</h1>"))
}

// Sending all the courses available in DataBase
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allCourses)
}

// Sending only one course from database
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)

	// Iterationg and checking if the course is available or not..
	for _, value := range allCourses {
		if value.CourseID == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}

	// If the course is no where to be found..
	json.NewEncoder(w).Encode("No such course is found")
}

// Creating a new course and injecting into DB
func createNewCourse(w http.ResponseWriter, r *http.Request) {
	// Before injecting check for some condition's
	if r.Body != nil {
		json.NewEncoder(w).Encode("Enter some Data")
	}

	// What is user sends empty data {}
	var course coureses
	_ = json.NewDecoder(r.Body).Decode(&course)
	if !course.validIDAndName() {
		json.NewEncoder(w).Encode("Add some data")
	}

	rand.Seed()
}
