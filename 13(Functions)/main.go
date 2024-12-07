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

	// Function that adds how many values as you like
	anotherResult := AnotherAdder(10, 20, 30, 40)
	fmt.Println("The result is:", anotherResult)

	// Return two different data types using Function's
	ans, message := returnTwoFunction(10, 20)
	fmt.Println("The result is:", ans)
	fmt.Println("The message is:", message)
}

func Namasthey() {
	fmt.Println("Namasthey Everyone") // Does not have any return type..
}

func Adder(a int, b int) int {
	return (a + b) // Returns the integer value..
}

func AnotherAdder(values ...int) int {
	// ... converts the inputs into a slice
	// we can pass as many values as we like
	// but with same datatype
	result := 0
	for _, val := range values {
		result += val
	}
	return result
}

// Function that returns sum(int) and a message(string)
// We can return multiple datatypes..
func returnTwoFunction(x int, y int) (int, string) {
	return x + y, "The numbers are successfully added"
}
