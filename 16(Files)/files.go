// In this understand about working with files in GO

// os - Create a file
// io - to write into it
// ioutils - to read the file (depricated)
// os.readfiles - to read the file

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Let's understand working with file's in GO language")

	// Creating
	file, err := os.Create("./file_operations.txt") //creating an txt file in current directory
	CheckNilError(err)                              // Checks for error

	// Writing into created file
	content := "let's GO" // Content to be written in File
	operationWrite, err := io.WriteString(file, content)
	CheckNilError(err)
	fmt.Println("Created the file successfully", operationWrite)
	// reding the content of the file

	reading, err := os.ReadFile("./file_operations.txt")
	fmt.Println("The message inside the file is:", string(reading)) // Give in Bites converting into string using string() method

}

// This function is used for checking the error
func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
