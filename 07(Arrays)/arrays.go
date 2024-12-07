// In this we will learn about Array's in Go-LaNg

package main

import "fmt"

func main() {
	fmt.Println("Let's learn about Arrays")
	// Declaration of an Array
	// var NameOfTheVariable [size]Datatype
	// var NameOfTheVariable = [size]Datatype{"0", "1", "2"}

	// fruit's list
	var myFruitList [3]string
	myFruitList[0] = "Apple"
	myFruitList[1] = "Banana"
	myFruitList[2] = "Grapes"
	fmt.Println("The Fruit list is : ", myFruitList)

	// Vegetables List
	// Direct Declaration
	var VegetablesList = [5]string{"Cabage", "MushRoom", "Potato"}
	fmt.Println("The Vegetables list is: ", VegetablesList)
	// # (important Point: If no value is initiliazed in Go language that is taken as "" in Go lang for string Type)
	// For int it might be 0,
	// For float it might be 0.0.

	// To find the length of an Array us len function
	fmt.Println("The Length of Vegetables list is: ", len(VegetablesList))

	var dummyData = [...]int{1, 2, 3, 4, 5}
	fmt.Println(dummyData)
	// In Go, [...] is a syntax used to define an array with a size inferred by the compiler based on the number of
	// elements provided during initialization
}
