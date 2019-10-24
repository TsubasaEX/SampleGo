// HelloWorld project main.go
package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch <- "Hello"
	}()

	go func() {
		ch2 <- "World"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
