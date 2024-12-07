// In this we are going to know how to Take user input's
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a value: ")
	value, _ := reader.ReadString('\n')
	fmt.Println("The Entered Value is:", value)
}
