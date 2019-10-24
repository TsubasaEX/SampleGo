// HelloWorld project main.go
package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.Tick(time.Millisecond * 100)
	bomb := time.After(time.Millisecond * 500)

	// The default case in a select is run if no other case is ready.
	// Use a default case to try a send or receive without blocking:
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-bomb:
			fmt.Println("bomb!!")
			return
		default:
			fmt.Println(".")
			time.Sleep(time.Millisecond * 50)
		}
	}
}
