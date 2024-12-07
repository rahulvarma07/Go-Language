// In this we will learn about methods
// Generally methods are function's in classes
// But in go we do not classes instead we have Strct's

package main

import "fmt"

// Let's understand struct by simple program
// Program := Check whether the user is active or not

// Creating a Struct
type UserInformation struct {
	Name           string
	IsUserActive   bool
	LastActiveDate string
}

// Method
func (user UserInformation) UserActiveDetails() {
	fmt.Println("Is user active: ", user.IsUserActive)
	fmt.Println("Last Active date: ", user.LastActiveDate)
}

// main
func main() {
	user1 := UserInformation{"User101", true, "07/12/2024"}
	user1.UserActiveDetails() // Calling the method..
}
