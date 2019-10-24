// HelloWorld project main.go
package main

import (
	"fmt"
	"sync"
)

type Account struct {
	pay  int
	name string
	mux  sync.Mutex
}

// func (a *Account) inc(){
// 	a.pay++
// }

func (a *Account) multi() {
	a.mux.Lock()
	a.pay += 1
	a.mux.Unlock()
}

func (a *Account) print() {
	a.mux.Lock()
	fmt.Println(a.pay)
	a.mux.Unlock()
}

func (a *Account) get() int {
	a.mux.Lock()
	defer a.mux.Unlock()
	return a.pay
}

func main() {
	a := Account{pay: 1, name: "zach"}
	for i := 0; i < 100; i++ {
		go a.multi()
		go a.print()
	}

	// time.Sleep(time.Second * 3)
	for {
		if a.get() == 101 {
			fmt.Println("final:", a.get())
			break
		}
	}
}
