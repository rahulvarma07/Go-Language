package main

import "fmt"

func main() {
	//fmt.Printf("This is Variables in gO LanG")
	// var keyword to declare variables
	// Syntax of Declaring a Variable is given as
	// var Variable_Name Type = value

	// 1. Declaration of Type String
	var nameOfTheUser string = "Rahul"
	fmt.Println(nameOfTheUser)
	fmt.Printf("The Type of the variable is: %T \n", nameOfTheUser)

	// 2. Declaration of Type bool
	var isUserLoggedIN bool = false
	fmt.Println(isUserLoggedIN)
	fmt.Printf("The Type of the Variable is: %T \n", isUserLoggedIN)

	// 3. Declaration of Type uint
	// -> Only has positive values
	// unit8, unit16, unit32, unit 64

	// 4. Declaration of Type int
	// -> Has negative and Positive integers
	// int8, int16, int32, int64

	// 5. Declaration of Type float
	// Decimal Point values
	// float32, float64

	var valueInRuppess float32 = 64.121212
	fmt.Println(valueInRuppess)
	fmt.Printf("The Type of the variable is: %T \n", valueInRuppess)

	// Default initialization
	// -> for string value it does not store anything
	// -> for int value it stors 0
	// -> for float value it stors 0
	var defaultData float32
	fmt.Println(defaultData)

	// Another way of declaration of a var
	// Without Specification of the data Type
	var anotherVariable = 1000
	fmt.Println(anotherVariable)

	// Another way of declaration without var Keyword
	// No var
	// Can't use this type of declaration in Gloabal'
	ageOftheUser := 12
	fmt.Println(ageOftheUser)

	// Const KeyWord
	const ConstantVariable int = 10
	fmt.Println(ConstantVariable)
}
