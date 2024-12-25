// In this let's understand how to build our own API's

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	fmt.Println("Creating the routes")

	// Creating a router
	rout := mux.NewRouter()

	// seeding..
	allCourses = append(allCourses,
		coureses{CourseName: "Go-Language", CourseID: "123", CourseFee: 0, CourseTutor: &courseTutor{TutorName: "Hithesh Choudary"}},
		coureses{CourseName: "Flutter-Language", CourseID: "124", CourseFee: 35000, CourseTutor: &courseTutor{TutorName: "Vasanth"}},
		coureses{CourseName: "DSA", CourseID: "125", CourseFee: 40000, CourseTutor: &courseTutor{TutorName: "Ashok"}},
	)

	// wrking with routes..

	// First page of home
	rout.HandleFunc("/", serveHome).Methods("GET")

	// Get all courses available
	rout.HandleFunc("/courses", getAllCourses).Methods("GET")

	// Get only one course based on the id
	rout.HandleFunc("/course/{id}", getOneCourse).Methods("GET")

	// Create a new course
	rout.HandleFunc("/addcourse", createNewCourse).Methods("POST")

	// update one course
	rout.HandleFunc("/updateonecourse", updateOneCourse).Methods("PUT")

	// delete one course
	rout.HandleFunc("delete/{id}", deleteACourse).Methods("DELETE")

	// Listening to a port
	log.Fatal(http.ListenAndServe(":7000", rout))
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

	// Creating a random course ID
	rand.NewSource(time.Now().UnixNano())
	// Converting into string beacuse CourseId is of type string
	randomId := strconv.Itoa(rand.Intn(10000))

	// Setting the courseID
	course.CourseID = randomId
	allCourses = append(allCourses, course)

	// Response
	json.NewEncoder(w).Encode(course)
}

// Updating Course controller's
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating a course data")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Update := (delete the exsisting course and add the new course)

	for ind, cour := range allCourses {
		if cour.CourseID == params["id"] {
			// for deleting a particular index in slice
			allCourses = append(allCourses[:ind], allCourses[ind+1:]...)

			var course coureses
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			allCourses = append(allCourses, course)
			json.NewEncoder(w).Encode(course)
		}
	}
}

// deleting a Course

func deleteACourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleting a course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for ind, cour := range allCourses {
		if cour.CourseID == params["id"] {
			allCourses = append(allCourses[:ind], allCourses[ind+1:]...)
			json.NewEncoder(w).Encode("Successfully deleted")
		}
	}
}
