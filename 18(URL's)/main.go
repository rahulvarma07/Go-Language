// Let's learn about URL's in GO language..
// Uniform resourse locater (URL)..
package main

import (
	"fmt"
	"net/url"
)

func main() {
	const Courseurl = "https://learning.dev:300/learn?coursename=GOlang&topic=Url"
	res, _ := url.Parse(Courseurl)

	// Let's check the type of res
	fmt.Printf("The type of res is: %T\n", res) // *url.URL (pointer to url.URL)
	// url.URL is a package of net/url that helps breaking down into part's

	// Let's see parts of the URL
	fmt.Println(res.Scheme)     // https
	fmt.Println(res.Host)       // learning.dev:300
	fmt.Println(res.Hostname()) // learning.dev
	fmt.Println(res.Port())     // 300
	fmt.Println(res.Path)       // learn
	fmt.Println(res.RawQuery)   // coursename=GOlang&topic=Url

	// Extracting RawQueries
	queryParams := res.Query()
	// Checking the type of queryParams
	fmt.Printf("The type of parameters is: %T\n", queryParams) // url.values
	// url.values simple means key value pair's also refered as Map's
	fmt.Println(queryParams)

	// Another important thing in the topic of URL
	// let's assume we have chunks of information about URL
	// and have to construct a complete url from it then we can do it as follows

	// ** Always pass a reference ** //
	newCourseUrl := &url.URL{
		Scheme:  "https",
		Host:    "LearnGO.dev",
		Path:    "GoLang",
		RawPath: "Topic=creatingUrl",
	}
	// To convert it into a string
	finURL := newCourseUrl.String()
	fmt.Println(finURL)
}
