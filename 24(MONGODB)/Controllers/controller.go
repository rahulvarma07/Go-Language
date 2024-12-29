// In This file we are going to write the logic's (CRUD) logics

package controllers

import (
	"context"
	"fmt"

	models "github.com/rahulvarma07/APIusingMONGO/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Make shore to keep name with starting with capital..
// To export..
type EmployeeRepo struct {
	Collection *mongo.Collection
}

// Make shore while writing method's Start the method name with capital letter's

// Method for Inserting a new employee..
func (r *EmployeeRepo) InsertNewEmployee(m *models.EmployeeDetails) (interface{}, error) {
	res, err := r.Collection.InsertOne(context.TODO(), m)

	// if there is an err
	if err != nil {
		return nil, err
	}

	// if there is no err
	return res.InsertedID, nil
}

// Method for finding an employye
// Accepts employee ID as an parameter
func (r *EmployeeRepo) FindEmployeeByID(empID string) (*models.EmployeeDetails, error) {
	// Creating a struct of employee ID
	var res models.EmployeeDetails

	// if success then reference with res
	err := r.Collection.FindOne(context.Background(), bson.M{"employee_id": empID}).Decode(&res)

	fmt.Println(res)

	// Classic err check..
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Method to find all employee..
// Does not require any parameter
func (r *EmployeeRepo) FindAllEmployee() ([]models.EmployeeDetails, error) {
	// Extracting all the collections in DB
	res, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var emp []models.EmployeeDetails
	// Passing emp as an reference and extracting all the collection's
	err = res.All(context.TODO(), &emp)

	// Classic err check
	if err != nil {
		return nil, err
	}

	return emp, err
}

// Update an employee record
func (r *EmployeeRepo) UpdateOneEmployye(id string, employeeDet *models.EmployeeDetails) (int64, error) {

	filter := bson.M{"employee_id": id}

	updated := bson.M{
		"$set": bson.M{
			"employee_id":         employeeDet.EmployeeID,
			"employee_name":       employeeDet.EmployeeName,
			"employee_department": employeeDet.EmployeeDepartment,
		},
	}

	res, err := r.Collection.UpdateOne(context.Background(),
		filter,  // finding the matching ID
		updated) // Set the employee details..
	// $set to update understood by mongo DB

	// Classic error check
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}

// Method to delete one employee...
func (r *EmployeeRepo) DeleteOneEmployee(id string) (int64, error) {
	res, err := r.Collection.DeleteOne(context.Background(), bson.M{"employee_id": id})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

// Func to delete all employee's from the DATA BASE
func (r *EmployeeRepo) DeleteAllEmployye() (int64, error) {
	res, err := r.Collection.DeleteMany(context.TODO(), bson.D{})

	// Classic error check
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
