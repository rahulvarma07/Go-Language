// let's learn about loop's

package main

import "fmt"

func main() {
	fmt.Println("Let's learn about loop's in golang")
	// declaring a slice to iterate through using for loop
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// type-1;
	// Based on index
	for i := 0; i < len(numbers); i++ {
		fmt.Println("The value at index ", i, " is equal to ", numbers[i])
	}

	// type-2;
	// Also based on index
	for i := range numbers {
		fmt.Println(numbers[i])
	}

	// Excersie try solving Two sum (LeetCode) Optimized solution using maps

	// for loop for direct iteration
	for _, elem := range numbers {
		fmt.Printf("The value are %v\n", elem)
	}

	// In go there is no such loop as a while loop
	// we can use for as a while

	loopTillTheCountIsTen := 1
	for loopTillTheCountIsTen <= 10 {
		fmt.Println(loopTillTheCountIsTen)
		loopTillTheCountIsTen++
	}

	// break := Terminates the loop
	// continue := Skips a-part in the loop
}
