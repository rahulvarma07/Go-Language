// In this we will be learning about function's in goLanguage
package main

import "fmt"

func main() {
	// Main always acts as an entry point..

	// Function that prints "namasthey everyone"
	Namasthey() // Calling the function

	// Function that adds two number's
	result := Adder(10, 20)
	fmt.Println("The result is:", result)

	anotherResult := AnotherAdder(10, 20, 30, 40)
	fmt.Println("The result is:", anotherResult)
}

func Namasthey() {
	fmt.Println("Namasthey Everyone") // Does not have any return type..
}

func Adder(a int, b int) int {
	return (a + b) // Returns the integer value..
}

func AnotherAdder(values ...int) int {
	result := 0
	for _, val := range values {
		fmt.Print(val, " ")
		result += val
	}
	return result
}
