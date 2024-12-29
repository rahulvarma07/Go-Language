package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	routers "github.com/rahulvarma07/APIusingMONGO/Routers"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const ConnectionString string = "mongodb+srv://rahulvarma070707:Db8DEVjpa0yWWloT@practice.dbxkl.mongodb.net/?retryWrites=true&w=majority&appName=Practice"
const DBName string = "Company_DataBase"
const CollectionName = "Employee_data"

var mongoClient *mongo.Client

func init() {
	// Building a connection
	//buildConnection, err := mongo.Connect(context.Background(), options.Client().ApplyURI(ConnectionString))
	mongoClient, err := mongo.Connect(options.Client().ApplyURI(ConnectionString))

	// Error check
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfull connection with DB")

	// Ping Checking..
	err = mongoClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ping Success")
}

func main() {
	defer mongoClient.Disconnect(context.Background())

	mongoColl := mongoClient.Database(DBName).Collection(CollectionName)

	dataColl := routers.EmployeeService{
		MongoCollection: mongoColl,
	}

	rout := mux.NewRouter()
	//rout.Handle("/", ServeHome).Methods("GET")
	rout.HandleFunc("/employee", dataColl.CreateNewEmployee).Methods("POST")
	rout.HandleFunc("/getemployee/{id}", dataColl.FindTheEmployeeByID).Methods("GET")
	rout.HandleFunc("/deleteemployee/{id}", dataColl.DeleteOneEmployeeByID).Methods("DELETE")
	rout.HandleFunc("/updateemployee/{id}", dataColl.UpdateOneEmployeeByID).Methods("PUT")

	log.Fatal(http.ListenAndServe(":7000", rout))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Building An API</h1>"))
}
