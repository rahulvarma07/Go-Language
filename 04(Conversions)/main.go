// In this file we will learn how to do conversion's of Data Type's
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a value: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	fmt.Println("The Value that has been Entered:", value)
	newValue, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		fmt.Println("There is an Error in Conversion: ", err)
	} else {
		fmt.Println("The value after conversion and Adding 1: ", newValue+1)
	}
}
