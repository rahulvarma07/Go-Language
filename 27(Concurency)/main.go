// In this we will learn about GO concurency..

// Concurency means performing multiple tasks simuntaneously..

/*
   Instead of performing one task at a time Go allows to push the
   task to the background using goRoutings
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // The number in Add() should be the number of routings..

	go func() {
		count1("sheep")
		wg.Done() // This Done decrements the counter by one everytime..
	}()

	count("Lion")

	wg.Wait() // Waits until all the background task's are completed..
}

func count(str string) {
	for i := 1; i <= 10; i++ {
		fmt.Println(str)
		time.Sleep(time.Millisecond * 500)
	}
}

func count1(str string) {
	for i := 1; i <= 20; i++ {
		fmt.Println(str)
		time.Sleep(time.Millisecond * 500)
	}
}
