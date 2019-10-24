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

//send-only type chan<- string
func ping(ch3 chan<- string, s string) {
	defer wg.Done()
	ch3 <- s
	//error
	// fmt.Println(<-ch3)
}

//receive-only type <-chan string
func pong(ch3 <-chan string) {
	defer wg.Done()
	fmt.Println(<-ch3)
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
		// time.Sleep(time.Second)
		wg2.Add(1)
		go get(ch2)
	}
	wg.Wait()
	wg2.Wait()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch)
	// }

	//A sender can close a channel to indicate that no more values will be sent.
	//Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	close(ch)
	close(ch2)

	//ok is false if there are no more values to receive and the channel is closed.
	v, ok := <-ch2
	fmt.Println(v, ok)

	//Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
	for item := range ch2 {
		fmt.Println(item)
	}

	ch3 := make(chan string, 2)
	wg.Add(2)
	go ping(ch3, "abc")
	go pong(ch3)
	wg.Wait()
	// fmt.Println(<-ch3)

}
