package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rahulvarma07/buildingAPI/routers"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var MongoClient *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("There is an issue with env")
		log.Fatal(err)
	}

	MongoClient, err = mongo.Connect(options.Client().ApplyURI(os.Getenv("connection_string")))

	if err != nil {
		log.Println("There is issue while connecting")
		log.Fatal(err)
	}

	err = MongoClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Println("Ping Unsuccessfull")
		log.Fatal(err)
	}

	log.Println("Ping Success")
}

func main() {

	defer MongoClient.Disconnect(context.Background())

	dataBase := MongoClient.Database(os.Getenv("database_name")).Collection(os.Getenv("database_collection"))

	ConnectionBase := routers.EmployeeData{
		Collection: dataBase,
	}

	rout := mux.NewRouter()

	rout.HandleFunc("/find-a-todo/{id}", ConnectionBase.FindTODOByID).Methods("GET")
	rout.HandleFunc("/find-a-todo", ConnectionBase.FindAllTODO).Methods("GET")
	rout.HandleFunc("/add-a-todo", ConnectionBase.InsertTODO).Methods("POST")
	rout.HandleFunc("/update-a-todo/{id}", ConnectionBase.UpdateATODO).Methods("PUT")
	// TODO: DELETE NOT WORKING CHECK IT
	rout.HandleFunc("/delete-todo/{id}", ConnectionBase.DeleteATODO).Methods("DELETE")
	rout.HandleFunc("/delete-all-todo", ConnectionBase.DeleteAllTODO).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", rout))
}
