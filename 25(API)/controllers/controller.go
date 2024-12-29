package controllers

import (
	"context"
	"log"

	"github.com/rahulvarma07/buildingAPI/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Created a mongo Collection
type ToDoData struct {
	Collection *mongo.Collection // I will pass when i create routers
}

// To Add a TODO to Database
func (tdd *ToDoData) AddAToDo(dataToInsert *models.ToDoModel) (bool, error) {
	res, err := tdd.Collection.InsertOne(context.Background(), dataToInsert)

	if err != nil {
		log.Println("#CONTROLLER_LOGIC : There is an error in inserting")
		return false, err
	}

	log.Println("Inserted a Todo with ID", res.InsertedID)
	return true, nil
}

// To find one TODO from dataBase
func (tdd *ToDoData) GetOneByID(todoID string) (*models.ToDoModel, error) {
	var res models.ToDoModel

	err := tdd.Collection.FindOne(context.Background(), bson.M{"todo_id": todoID}).Decode(&res)

	if err != nil {
		log.Println("#CONTROLLER_LOGIC There is an error")
		return nil, err
	}

	log.Println("Successfully found")
	return &res, nil
}

// To find all the TODO from the database
func (tdd *ToDoData) FindAll() ([]models.ToDoModel, error) {
	var result []models.ToDoModel
	res, err := tdd.Collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Println("#CONTROLLER_LOGIC there is an error in getting all the todo's")
		return nil, err
	}

	err = res.All(context.Background(), &result)

	if err != nil {
		log.Println("#CONTROLLER_LOGIC there is an error in getting all the todo's")
		return nil, err
	}

	log.Println("Successfully Got all the todo's")

	return result, nil
}

// To update one todo
func (tdd *ToDoData) UpdateANote(ToDoId string, dataToBeUpdated *models.ToDoModel) (bool, error) {

	update := bson.M{
		"$set": bson.M{
			"task":           dataToBeUpdated.Task,
			"task_completed": dataToBeUpdated.TaskCompleted,
		},
	}

	res, err := tdd.Collection.UpdateOne(context.Background(), bson.M{"todo_id": ToDoId}, update)

	if err != nil {
		log.Println("#CONTROLLER_LOGIC there is an error in updateing..")
		return false, err
	}

	log.Println("Successfully updated the todo #UpdateCount := ", res.ModifiedCount)
	return true, nil
}

// To dlete one TODO from the Database
func (tdd *ToDoData) DeleteOneTODO(ToDoId string) (bool, error) {
	res, err := tdd.Collection.DeleteOne(context.Background(), bson.M{"todo_id": ToDoId})

	if err != nil {
		log.Println("#CONTROLLER_LOGIC there is an error in deleteing")
		return false, err
	}

	log.Println("Successfully deleted the TODO", res.DeletedCount)
	return true, nil
}

// To delete the entire data of TODO's from DataBase
func (tdd *ToDoData) DeleteAll() (bool, error) {
	res, err := tdd.Collection.DeleteMany(context.Background(), bson.M{})

	if err != nil {
		log.Println("#CONTROLLER_LOGIC there is an error in deleteing everything")
		return false, err
	}

	log.Println("Successfully Deleted, DeletedCount := ", res.DeletedCount)
	return true, nil
}
