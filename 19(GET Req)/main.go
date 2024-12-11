// In this let's understand GET request's

package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Performing GET Request's using GO")
	GetRequests()
}

func GetRequests() {
	const getUrl string = "https://catfact.ninja/fact"
	// The above string is an public API for random Cat Fact's

	response, err := http.Get(getUrl) // Get method using http package
	checkNilError(err)

	// Checking some important parameters while making a Get Request
	fmt.Println(response.StatusCode)    // Give's the status code of the responce
	fmt.Println(response.ContentLength) // Give's the length of the content
	// There are few more try exploring it..

	defer response.Body.Close() // Devoloper's Responsibility to Close the
	// Request once the task is done

	// Reading the responce
	content, err := io.ReadAll(response.Body)
	checkNilError(err)

	// Initially the content is of Byte datatype

	// A simpler way of converting the byte data into string
	// finalResponce := string(content)
	// fmt.Println(finalResponce)

	// Another way using strings package..
	// This package helps us to have more hold on the raw data
	// If we see we are converting the above byte data into strings directly hence we cannot have the byte data back
	// So to overcome this we can use strings package provided by GO

	var responseString strings.Builder
	//byteResponse, _ := responseString.Write(content)
	responseString.Write(content)
	//fmt.Println(byteResponse) // Gives the byte count
	fmt.Println(responseString.String())
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
