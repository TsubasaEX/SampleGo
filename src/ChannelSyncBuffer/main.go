// HelloWorld project main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var wg2 sync.WaitGroup

func plus(ch chan int, a int) {
	defer wg.Done()
	ch <- (a + 10)
}

func multi(ch2 chan int, ch chan int) {
	defer wg.Done()
	ch2 <- (<-ch) * (-1)
}

func get(ch2 chan int) {
	defer wg2.Done()
	fmt.Println(<-ch2)
}

func main() {

	//define buffer len equal to or more than the num of routines
	ch := make(chan int, 10)
	ch2 := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go plus(ch, i)
	}
	wg.Wait()

	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go multi(ch2, ch)
		time.Sleep(time.Second)
		wg2.Add(1)
		go get(ch2)
	}
	wg.Wait()
	wg2.Wait()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch)
	// }
	close(ch)
	close(ch2)
	for item := range ch2 {
		fmt.Println(item)
	}

}
