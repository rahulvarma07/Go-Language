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
	createJSON()
}

func createJSON() {

	// Creating slice's of type's structs(studentPortalDetais)..
	normalData := []studentPortalDetais{
		{"Bob", "CS", 7.1, "ecap", []string{"CS-1", "CS-2", "CS-3"}},
		{"Alice", "CSE", 9.1, "ecap", []string{"CSE-1", "CSE-2", "CSE-3"}},
		{"Abram", "ECE", 8.1, "ecap", []string{"ECE-1", "ECE-2", "ECE-3"}},
		{"Tate", "EEE", 8.1, "ecap", []string{"EEE-1", "EEE-2", "EEE-3"}},
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
}
