// HelloWorld project main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

// Goroutines
func main() {
	// say("Hello")
	// say("World")

	// go say("Hello")
	// say("World")

	// nothing happend
	// go say("Hello")
	// go say("World")

	// say("Hello")
	// go say("World")

	// bad idea
	// go say("Hello")
	// go say("World")
	// time.Sleep(time.Second * 2)

	//Sync Concurrency
	// nothing happend
	// wg.Add(1)
	// go say("Hello")
	// wg.Add(-1)
	// go say("World")
	// wg.Wait()

	// fatal error
	// wg.Add(1)
	// go say("Hello")
	// wg.Add(-2)
	// go say("World")
	// wg.Wait()

	// good
	// wg.Add(2)
	// go say("Hello")
	// go say("World")
	// wg.Wait()

	// good
	wg.Add(1)
	go say("Hello")
	wg.Add(1)
	go say("World")
	wg.Wait()

	wg.Add(1)
	go say("There")
	wg.Wait()
}
