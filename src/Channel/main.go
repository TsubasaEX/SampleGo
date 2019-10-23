// HelloWorld project main.go
package main

import (
	"fmt"
)

func say(c chan string, s string) {
	c <- s
}

func main() {

	ch := make(chan string)

	go say(ch, "Hello")
	go say(ch, "World")

	s1 := <-ch
	s2 := <-ch

	fmt.Println(s1)
	fmt.Println(s2)

}
