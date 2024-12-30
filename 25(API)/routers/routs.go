package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rahulvarma07/buildingAPI/controllers"
	"github.com/rahulvarma07/buildingAPI/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EmployeeData struct {
	Collection *mongo.Collection
}

type Res struct {
	Data                 string `json:"data_received" bson:"data_received"`
	ReqType              string `json:"request_type" bson:"request_type"`
	OperationSuccessfull bool   `json:"status" bson:"status"`
	Err                  error  `json:"err" bson:"err"`
}

type Res1 struct {
	Data interface{} `json:"data_received" bson:"data_received"`
	Err  error
}

// Inserting a New employee
func (ed *EmployeeData) InsertTODO(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Res{}

	defer r.Body.Close()
	defer json.NewEncoder(w).Encode(res)

	newID := uuid.NewString()

	var EmpToBeInserted models.ToDoModel

	err := json.NewDecoder(r.Body).Decode(&EmpToBeInserted)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("#ROUTER_ERROR There is an error in adding a new TODO")
		res.Data = ""
		res.ReqType = "POST"
		res.OperationSuccessfull = false
		res.Err = err
		return
	}

	EmpToBeInserted.ToDoID = newID
	EmpToBeInserted.TimeStamp = time.DateTime
	EmpToBeInserted.TaskCompleted = false

	dataCollection := controllers.ToDoData{
		Collection: ed.Collection,
	}

	response, err := dataCollection.AddAToDo(&EmpToBeInserted)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("#ROUTER_ERROR There is an error in adding a new TODO to the database")
		res.Data = "NODATA"
		res.ReqType = "POST"
		res.OperationSuccessfull = response
		res.Err = err
		return
	}

	log.Println("Successfully added a new TODO to the database")
	w.WriteHeader(http.StatusAccepted)

	res.Data = newID
	res.OperationSuccessfull = response
	res.Err = nil
	res.ReqType = "POST"

}

func (ed *EmployeeData) UpdateATODO(w http.ResponseWriter, r *http.Request) {
	res := &Res{}

	w.Header().Add("Content-Type", "application/json")

	defer r.Body.Close()
	defer json.NewEncoder(w).Encode(res)

	var EmpToBeUpdated models.ToDoModel

	params := mux.Vars(r)["id"] // ID to be updated

	dataCollection := controllers.ToDoData{
		Collection: ed.Collection,
	}

	error := json.NewDecoder(r.Body).Decode(&EmpToBeUpdated)

	if error != nil {
		log.Println("Bad request type's")
		w.WriteHeader(http.StatusBadRequest)
		res.Data = "No data due to bad request"
		res.Err = error
		res.OperationSuccessfull = false
		res.ReqType = "PUT"
		return
	}

	response, err := dataCollection.UpdateANote(params, &EmpToBeUpdated)

	if err != nil {
		log.Println("Bad request type's")
		w.WriteHeader(http.StatusBadRequest)
		res.Data = "No data due to bad request"
		res.Err = err
		res.OperationSuccessfull = response
		res.ReqType = "PUT"
		return
	}

	log.Println("Successfulyy updated", EmpToBeUpdated.ToDoID)
	res.Data = EmpToBeUpdated.Task
	res.OperationSuccessfull = response
	res.ReqType = "PUT"
	res.Err = nil
}

func (ed *EmployeeData) FindTODOByID(w http.ResponseWriter, r *http.Request) {
	res := &Res1{}

	w.Header().Add("Content-Type", "application/json")

	defer r.Body.Close()
	defer json.NewEncoder(w).Encode(res)

	params := mux.Vars(r)["id"]

	dataCollection := controllers.ToDoData{
		Collection: ed.Collection,
	}

	response, err := dataCollection.GetOneByID(params)

	if err != nil {
		log.Println("Bad request type's")
		w.WriteHeader(http.StatusBadRequest)
		res.Data = "No data due to bad request"
		res.Err = err
		return
	}

	log.Println("Successfully found the TODO note")

	w.WriteHeader(http.StatusAccepted)
	res.Data = response
	res.Err = nil
}

func (ed *EmployeeData) FindAllTODO(w http.ResponseWriter, r *http.Request) {
	res := &Res1{}

	w.Header().Add("Content-Type", "application/json")

	defer r.Body.Close()
	defer json.NewEncoder(w).Encode(res)

	dataCollection := controllers.ToDoData{
		Collection: ed.Collection,
	}

	response, err := dataCollection.FindAll()

	if err != nil {
		log.Println("Bad request type's")
		w.WriteHeader(http.StatusBadRequest)
		res.Data = "No data due to bad request"
		res.Err = err
	}

	log.Println("Successfully found the TODO note")

	w.WriteHeader(http.StatusAccepted)
	res.Data = response
	res.Err = nil
}

func (ed *EmployeeData) DeleteATODO(w http.ResponseWriter, r *http.Request) {
	res := &Res{}

	w.Header().Add("Content-Type", "application/json")

	defer r.Body.Close()
	defer json.NewEncoder(w).Encode(res)

	params := mux.Vars(r)["id"]

	dataCollection := controllers.ToDoData{
		Collection: ed.Collection,
	}

	response, err := dataCollection.DeleteOneTODO(params)

	if err != nil {
		log.Println("Bad request type's")
		w.WriteHeader(http.StatusBadRequest)
		res.Data = "No data due to bad request"
		res.Err = err
		res.OperationSuccessfull = response
		res.ReqType = "DELETE"
		return
	}

	log.Println("Successfully Deleted")
	w.WriteHeader(http.StatusAccepted)
	res.Data = "Success"
	res.Err = nil
	res.OperationSuccessfull = response
	res.ReqType = "DELETE"

}

func (ed *EmployeeData) DeleteAllTODO(w http.ResponseWriter, r *http.Request) {
	res := &Res{}

	w.Header().Add("Content-Type", "application/json")

	defer r.Body.Close()
	defer json.NewEncoder(w).Encode(res)

	dataCollection := controllers.ToDoData{
		Collection: ed.Collection,
	}

	response, err := dataCollection.DeleteAll()

	if err != nil {
		log.Println("Bad request type's")
		w.WriteHeader(http.StatusBadRequest)
		res.Data = "No data due to bad request"
		res.Err = err
		res.OperationSuccessfull = response
		res.ReqType = "DELETE"
		return
	}

	log.Println("Successfully Deleted")
	w.WriteHeader(http.StatusAccepted)
	res.Data = "Success"
	res.Err = nil
	res.OperationSuccessfull = response
	res.ReqType = "DELETE"
}
