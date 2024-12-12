// In this let's see how to make Post Request's using GO language..

package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Making Post Request's")
	PostRequests()
}

func PostRequests() {
	const attendenceAPI string = "Example.go"
	postData := strings.NewReader(`{"rool_no" : "22P31A04H1"}`) // Dont forget to use ` ` instead of ' '

	response, err := http.Post(attendenceAPI, "application/json", postData)
	checkNilError(err)

	defer response.Body.Close() // Closing once the task is done

	fmt.Println(response.StatusCode) // Checking the statusCode of the responce

	content, err := io.ReadAll(response.Body)

	fmt.Println(string(content)) // Printing the responce

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
