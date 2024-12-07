// cfs := Control flow statement's
// In this we will learn about control flow statement's
// if, else if, else
// switch case

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Let's learn about control statement's in Golang")

	// Let's create a program to check whether the user input is gretar, lesser or equla's to 10

	// Taking input from user
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	finalUserInput, _ := strconv.Atoi(strings.TrimSpace(userInput)) // Converting it into an integer

	// if, else if, else
	if finalUserInput < 10 {
		fmt.Println("Less Than 10")
	} else if finalUserInput > 10 {
		fmt.Println("Greater than 10")
	} else {
		fmt.Println("Equal's to 10")
	}

	// Another weird syntax in golang
	if num := 10; num < 10 { // In this line we intialized a variable to 10 as well as we checked for the condition
		fmt.Println("Num is less than 10")
	} else if num > 10 {
		fmt.Println("Num is greater than 10")
	} else {
		fmt.Println("Num is equal's 10")
	}

	// Now let's understand about switch case in GoLang
	// let's write a switch case program to check what day it is in the week
	// Example 1 := Monday, 7 := sunday
	fmt.Println("The below is an example program for Switch Case")
	dayOfTheWeek := 6
	switch dayOfTheWeek {
	case 1:
		fmt.Println("The day is Monday")
	case 2:
		fmt.Println("The day is Tuesday")
	case 3:
		fmt.Println("The day is Wednesday")
	case 4:
		fmt.Println("The day is Thursday")
		fallthrough // using an fallthrough mean's if this case is executed the below case will also be executes
	case 5:
		fmt.Println("The day is Friday")
	case 6:
		fmt.Println("The day is Saturday")
	case 7:
		fmt.Println("The day is Sunday")
	}
}
