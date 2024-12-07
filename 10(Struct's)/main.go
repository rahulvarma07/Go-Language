// Let's learn some basic's of Struct
// 1. Struct are like classe's in GO we dont have classes we have struct

package main

import "fmt"

// Declaration of struct
// type NameOfTheStructure struct{}
// NameOfTheStruct must start with a capital letter
type UserDetails struct {
	Name   string
	Email  string
	age    int
	IsMale bool
}

func main() {
	fmt.Println("Struct's in Go-laNg")
	idDetailsOfUser1 := UserDetails{"Rahul Varma", "rahul@learn.in", 20, true} // initializing

	// Printing the details in basic manner
	fmt.Println(idDetailsOfUser1)

	// Using printf ans %+v give's cleaner way as an output
	fmt.Printf("%+v\n", idDetailsOfUser1)

	// To only print required value from the idDetailsOfUser1
	fmt.Printf("%+v", idDetailsOfUser1.Name)
}
