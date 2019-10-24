// HelloWorld project main.go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	ch := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string, 3)

	go func() {
		ch <- "Hello"
	}()

	go func() {
		ch2 <- "World"
	}()

	go func() {
		for i := 0; i < 3; i++ {
			ch3 <- "fuck-" + strconv.Itoa(i)
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

	for i := 0; i < 3; i++ {
		select {
		case msg3 := <-ch3:
			fmt.Println(msg3)
		}
	}
}
