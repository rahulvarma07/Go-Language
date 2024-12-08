// In this we will learn about defer keyword
// In is basically used to manupliate the flow of the program
// what ever we write with defer it is executed at the very last of the program
// If multiple statements are there it follow's LIFO (last in frst out)

package main

import "fmt"

func main() {
	fmt.Println("Let's learn about defer in GO")
	// check the order of printing
	defer fmt.Println("G")
	defer fmt.Println("A")
	defer fmt.Println("N")
	defer fmt.Println("L")
	defer fmt.Println("-")
	defer fmt.Println("O")
	defer fmt.Println("G")
	// LIFO order is printed
}

// in the future we will understand the use cases of defer in GoLang
