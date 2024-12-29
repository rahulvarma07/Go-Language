package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	controllers "github.com/rahulvarma07/APIusingMONGO/Controllers"
	models "github.com/rahulvarma07/APIusingMONGO/Models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type EmpData struct {
	Data  interface{}
	Error string
}

func (ed *EmployeeService) CreateNewEmployee(w http.ResponseWriter, r *http.Request) {
	res := &EmpData{}

	w.Header().Add("Content-Type", "application/json")

	//defer r.Body.Close()                 // Closing the body at the very end
	defer json.NewEncoder(w).Encode(res) // Sending the data at very end

	// Creating a model Schema
	var empData models.EmployeeDetails

	// Putting the response budy inside empData after passing it's reference
	err := json.NewDecoder(r.Body).Decode(&empData)

	// Classic error checking
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Bad Request"
		return
	}

	// Creating a new UUID
	empId := uuid.NewString()

	// Calling the logic's
	empRepo := controllers.EmployeeRepo{
		Collection: ed.MongoCollection,
	}

	// Assining the UID created
	empData.EmployeeID = empId

	// Calling the method insertNewEmployee from contrllers
	response, err := empRepo.InsertNewEmployee(&empData)

	// Classic err checking
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	log.Println("Employee Successfully inserted with, ", response)
}

func (ed *EmployeeService) FindTheEmployeeByID(w http.ResponseWriter, r *http.Request) {
	// Setting the header
	w.Header().Add("Content-Type", "application/json")

	// Creating a res struct
	res := &EmpData{}

	//defer r.Body.Close()                 // Closing the res body once the task is done
	defer json.NewEncoder(w).Encode(res) // encoding and sending the response..

	// Taking the id into
	params := mux.Vars(r)["id"]

	// Creaing a empRepo from package controller's
	empRepo := controllers.EmployeeRepo{
		Collection: ed.MongoCollection,
	}

	response, err := empRepo.FindEmployeeByID(params)

	// Classic error checking
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Bad Request"
		return
	}

	res.Data = response
	w.WriteHeader(http.StatusAccepted)
	log.Println("Good response", response.EmployeeID)
}

func (ed *EmployeeService) DeleteOneEmployeeByID(w http.ResponseWriter, r *http.Request) {
	// Setting header's
	w.Header().Add("Content-Type", "application/json")

	res := &EmpData{}

	//defer r.Body.Close()
	defer json.NewEncoder(w).Encode(res)

	empRepo := controllers.EmployeeRepo{
		Collection: ed.MongoCollection,
	}

	params := mux.Vars(r)["id"]

	result, err := empRepo.DeleteOneEmployee(params)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad Request", err)
		res.Error = "Bad Request"
		return
	}

	res.Data = result
	w.WriteHeader(http.StatusAccepted)
	log.Println("Successfully deleted the employee")
}

func (ed *EmployeeService) UpdateOneEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &EmpData{}

	//defer r.Body.Close()
	defer json.NewEncoder(w).Encode(res)

	params := mux.Vars(r)["id"]

	var empData models.EmployeeDetails
	err := json.NewDecoder(r.Body).Decode(&empData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Bad Request"
		return
	}

	empData.EmployeeID = params

	empRepo := controllers.EmployeeRepo{
		Collection: ed.MongoCollection,
	}

	result, err := empRepo.UpdateOneEmployye(empData.EmployeeID, &empData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Error = "Bad Request"
		return
	}

	res.Data = result
	w.WriteHeader(http.StatusAccepted)
	log.Println("Updated employee ID", result)
}
