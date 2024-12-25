// Let's understand creating JSON data

package main

import (
	"encoding/json"
	"fmt"
)

type studentPortalDetais struct {
	Name           string `json:"nameOfTheStudent"`   // here by doing this we can manuplate the name's when converted into JSON
	Branch         string `json:"branchOfTheStudent"` // Instead of printing Branch in the output it give's branchOfTheStudent
	CGPA           float32
	Password       string   `json:"-"`                         // The password will not be visible
	CurrentSubject []string `json:"currentSubjects,omitempty"` // omitempty means if currentSubjects are null it does not display it.
}

func main() {
	fmt.Println("Let's Understand about creating JSON data")
	//createJSON()
	JsonDecode()
}

func createJSON() {

	// Creating slice's of type's structs(studentPortalDetais)..
	normalData := []studentPortalDetais{
		{"Bob", "CS", 7.1, "ecap", []string{"CS-1", "CS-2", "CS-3"}},
		{"Alice", "CSE", 9.1, "ecap", []string{"CSE-1", "CSE-2", "CSE-3"}},
		{"Abram", "ECE", 8.1, "ecap", []string{"ECE-1", "ECE-2", "ECE-3"}},
		{"Tate", "EEE", 8.1, "ecap", nil},
	}

	// To Convert the above data into JSON
	// Check the output to understand difference btwn Marshal & MarshalIndent
	// finalJsonData, err := json.Marshal(normalData)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(finalJsonData))

	finalJsonData, err := json.MarshalIndent(normalData, "", "\t")
	// Expects three parameters
	// 1.Data , 2.Prefix, 3.Indent
	// \t represents a tab character
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJsonData))

	// Cross check the output to understand this..
}

func JsonDecode() {

	jsonDataFromCall := []byte(`
	{
        "nameOfTheStudent": "Tate",
        "branchOfTheStudent": "EEE",
        "CGPA": 8.1
    }
	`)

	// If we want to check that if our JSON data coming from web or any other req is valid or not
	isValid := json.Valid(jsonDataFromCall)
	if isValid {
		fmt.Println("The JSON data is valid")
	} else {
		fmt.Println("THE JSON DATA IS NOT VALID")
		return
	}

	// Storing JSON data coming into a strct
	var StudentDetails1 studentPortalDetais
	// Let's Say jsonDataFromCall(is the received data from call's) & we want to store it in a struct
	json.Unmarshal(jsonDataFromCall, &StudentDetails1) // Store's it in a struct
	// Remember to pass as refernce to avoid copy's
	fmt.Printf("%#v\n", StudentDetails1)

	// Now let's say you dont alway's want to store it in a struct
	// you want to store it in a (Map){key : value}
	var receivedData map[string]interface{}
	json.Unmarshal(jsonDataFromCall, &receivedData)
	fmt.Printf("%#v\n", receivedData)

	// Access or loop throught the map's
	// for key, val := range receivedData {
	// 	fmt.Printf("The Key is %v & the Val is %v \n", key, val)
	// }
}
