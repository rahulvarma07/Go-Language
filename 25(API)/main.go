package main

import (
	"context"
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	rout := mux.NewRouter()

	// TODO: Add routs and write routers login
}
