// In This we will learn about pointer's
package main

import "fmt"

func main() {
	// Pointers are denoted by *
	// Directly connected with the address so no copy's are created

	// var myPtr *int
	// fmt.Println("The value of my pointer is: ", myPtr)

	// & symbol is used to reference the address of the value

	myPtr := 10
	var AddressOfMyPtr = &myPtr
	fmt.Println("The address of the variable myPtr is: ", AddressOfMyPtr) // & used to get the address of the variable
	fmt.Println("The value of the variable myPtr is: ", *AddressOfMyPtr)  // * is used for de-referencing the variable
}
