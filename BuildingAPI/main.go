// In this let's understand how to build our own API's

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Initially we did not work with DataBase
// So lets assume these are model(structure) of our database..
type coureses struct {
	CourseName  string       `json:"coursename"`
	CourseID    string       `json:"course_id"`
	CourseFee   float32      `json:"course_fee"`
	CourseTutor *courseTutor `json:"course_tutor"`
}

type courseTutor struct {
	TutorName string `json:"tutor_name"`
}

// Fake DataBase
var allCourses []coureses // Storing the model's inside a struct..

// Let's define some helper function
// helper function (used to valid the function)

// This function helps for checking whethre the course ID and
// name are not empty..

func (course *coureses) validIDAndName() bool {
	return (course.CourseName != "" && course.CourseID != "")
}

func main() {

}

// When Opened a local host ensures there is some HTML code
// Contrller's
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Building An API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Course")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(allCourses)
}
