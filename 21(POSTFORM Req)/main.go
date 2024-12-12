// In this lets understand about POST FORM Request's
// In POST there is another type ie, POST FORM
// Often used for file upload's

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Let's understand how to do POSTFORM Request's")
	postFormReq()
}

func postFormReq() {
	const DummyUrl string = "https://example.com/PostForm" // This is not valid URL
	// It's just an example

	data := url.Values{}
	data.Add("name", "Naruto")

	responce, err := http.PostForm(DummyUrl, data)
	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	finalData, _ := io.ReadAll(responce.Body) // ignoring the err

	var myfinalDatainString strings.Builder
	myfinalDatainString.Write(finalData)

	fmt.Println(myfinalDatainString.String())
}
