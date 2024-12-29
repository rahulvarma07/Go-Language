// Let's see the TESTING OF API(LOGOC) That we created
package routers_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	controllers "github.com/rahulvarma07/APIusingMONGO/Controllers"
	models "github.com/rahulvarma07/APIusingMONGO/Models"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://rahulvarma070707:Db8DEVjpa0yWWloT@practice.dbxkl.mongodb.net/?retryWrites=true&w=majority&appName=Practice"
const dbName = "Companay"
const dbCollectionName = "Employee_testing"

func connectToTheClient() *mongo.Client {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection success")

	err = mongoClient.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ping success")

	return mongoClient
}

func TestOpt(t *testing.T) {
	mongoTestCLient := connectToTheClient()                // function
	defer mongoTestCLient.Disconnect(context.Background()) // Disconnect once the use is done..

	// Creating a emplpoyee id
	empID1 := uuid.New().String()
	fmt.Println(empID1)

	// Creating a new collection in Data base
	newCollectionInDB := mongoTestCLient.Database(dbName).Collection(dbCollectionName)

	newRepo := controllers.EmployeeRepo{
		Collection: newCollectionInDB,
	}

	// Testing for inserting an employee
	t.Run("Insert one employee", func(t *testing.T) {
		// new schema of user..
		empDet := models.EmployeeDetails{
			EmployeeName:       "Alex",
			EmployeeID:         empID1,
			EmployeeDepartment: "Science",
		}

		res, err := newRepo.InsertNewEmployee(&empDet)

		// error checking..
		if err != nil {
			t.Fatal("inserting employee details unsuccessfull")
		}

		t.Log("Successfully inserted employee", res)
	})

	t.Run("Get employee by ID", func(t *testing.T) {
		// Checking for employee...
		empID := "1f92e483-0763-4a2a-93da-72fc9b6c27df"

		res, err := newRepo.FindEmployeeByID(empID)

		if err != nil {
			log.Fatal(err)
		}

		t.Log("Found the emp \n", res.EmployeeName)
	})

	t.Run("Update one employee", func(t *testing.T) {
		empID := "1f92e483-0763-4a2a-93da-72fc9b6c27df"
		updatedEmplDet := models.EmployeeDetails{
			EmployeeName:       "not Alex",
			EmployeeID:         "123",
			EmployeeDepartment: "Not Science",
		}

		res, err := newRepo.UpdateOneEmployye(empID, &updatedEmplDet)
		if err != nil {
			log.Fatal("Failed to update", err)
		}

		t.Log(res)
	})

	t.Run("Delete One Employee", func(t *testing.T) {
		uid := "123"
		res, err := newRepo.DeleteOneEmployee(uid)

		// Classic err checking
		if err != nil {
			log.Fatal("Failed to delete", err)
		}

		log.Println(res)
	})

	// t.Run("Find all employee's", func(t *testing.T) {
	// 	res, err := newRepo.FindAllEmployee()

	// 	if err != nil {
	// 		t.Fatal("No employee's", err)
	// 	}

	// 	t.Log("The employee's data \n", res)
	// })

}

// For testing use this command
// go test ./... -v
// ./ current directory
// ... checks for testing
// -v for checking the logs

// Make shore to save your file as name_test.go
// then only it will understand that this is a testing file...
