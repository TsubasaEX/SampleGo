// HelloWorld project main.go
package main

import (
	"errors"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("cleanup:", r)
	}
	wg.Done()
}

//pass the panic to cleanup
func cleanup2() {
	defer cleanup()
	if r := recover(); r != nil {
		fmt.Println("cleanup2:", r)
		panic(r)
	}
}

// pass the panic or the cleanup won't receive panic

// defer is FILO
func say(s string) {
	defer cleanup2()
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		if i == 3 {
			panic("i is 3")
		}
	}
}

func testerr() (string, error) {
	// return "no error", nil
	return "error", errors.New("myError")
}
func main() {
	wg.Add(1)
	go say("Hello")
	wg.Add(1)
	go say("World")
	wg.Wait()

	wg.Add(1)
	go func() {
		str, err := testerr()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(str)
		}
		wg.Done()
	}()
	wg.Wait()
}
