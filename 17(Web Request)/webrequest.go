// Let's learn about web request's in GO

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web Request's using GO lang")

	url := "https://maya.technicalhub.io/"
	responce, err := http.Get(url)
	CheckNilError(err)
	fmt.Printf("The Type of Response %T\n", responce)
	defer responce.Body.Close() // Caller's Responsilbility...

	body, err := io.ReadAll(responce.Body)
	CheckNilError(err)
	content := string(body)
	fmt.Println(content)
}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}
