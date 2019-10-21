// HelloWorld project main.go
package main

import (
	"fmt"
	"time"
)

func say(s string) {

	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}

}

func main() {
	// say("Hello")
	// say("World")

	// go say("Hello")
	// say("World")

	// go say("Hello")
	// go say("World")

	// say("Hello")
	// go say("World")

	go say("Hello")
	go say("World")
	time.Sleep(time.Second * 2)
}
