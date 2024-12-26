package controllers

import "go.mongodb.org/mongo-driver/mongo"

type employeeRepo struct {
	Collection *mongo.Collection
}
