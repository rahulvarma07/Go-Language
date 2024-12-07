// Slice's are like vector's in CPP
// Doe's not have a fixed length..
// Can append how many values as we want after declaration
// has similar syntax of Array's

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slic:s in Go-laNg")
	var myFruitList = []string{}
	// To add values to myFruitList use the method append.
	myFruitList = append(myFruitList, "Apple", "Banana", "Grapes")
	// What Append actually does is it creates modified list and updates to exsisting list..
	fmt.Println(myFruitList)

	// Slice a List we use the symbol ':'
	// Inclusive range..
	myFruitList = append(myFruitList[1:2], "Hey")
	fmt.Println("The list after using the slice operation is: ", myFruitList)

	// lets usderstand how to delete a particular index value in Slic:s
	var myVegeList = []string{"Tamato", "Carrot", "Spinach", "Bottle guard"}
	fmt.Println("The list before removing the carrot was :", myVegeList)
	myVegeList = append(myVegeList[:1], myVegeList[2:]...)
	fmt.Println("The list after removing the carrot was :", myVegeList)

	// sort.Strings(myVegeList) (used to sort the slices of string's)
}
