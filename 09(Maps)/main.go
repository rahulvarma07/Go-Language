// Maps are also called Hash Table's
// Key-Value pairs

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's learn about Map's in Go-laNg")
	// Map declaration syntax
	// NameOfTheMap := make(map["Key_DataType"]"value_DataType")

	// Let's create a map of Programming Languages
	// Key are the name of language
	// value's are it's Short form

	programmingLanguage := make(map[string]string)
	programmingLanguage["Java Script"] = "JS"
	programmingLanguage["Ruby"] = "RB"
	programmingLanguage["Python"] = "PY"

	fmt.Println(programmingLanguage)

	// to access from the Map
	// nameOfTheMap["Key_Value"]
	fmt.Println("The Short form of Python is: ", programmingLanguage["Python"])

	// Let's see how to Remove/Delete something from a map
	// syntax is
	// delete(nameOfTheMap, "key_value")
	fmt.Println("The map before deleting Ruby is: ", programmingLanguage)
	delete(programmingLanguage, "Ruby")
	fmt.Println("The map after deleting Ruby is: ", programmingLanguage)

	// To Check whether the language exist's in Map
	value, found := programmingLanguage["Python"]
	// found is true(IF EXSIST'S) false(IF NOT EXSIST'S)
	if found {
		fmt.Println("Found")
	} else {
		fmt.Println(value, "is not found")
	}

	// Another way
	if programmingLanguage["Ruby"] != "" {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}

	// Let's see how to loop in case of Map's
	for key, value := range programmingLanguage {
		fmt.Println("The key is: ", key, "and it's value is: ", value)
	}
}
