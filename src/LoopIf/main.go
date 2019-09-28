// LoopIf project main.go
package main

import (
	"fmt"
	"time"
)

func sum(a int, b int) int {
	var sum int
	for i := a; i <= b; i++ {
		sum += i
	}
	return sum
}

func definiteLoop(in *int) {
	var count int
	for count <= *in {
		fmt.Println(count)
		count++
		time.Sleep(time.Second)
	}
}

func infiniteLoop() {
	var count int
	time.Sleep(time.Second)
	for {
		fmt.Println(count)
		count++
		time.Sleep(time.Second)
		// if count++; count > 5 {
		// 	break
		// } else if count > 3 {
		// 	break
		// }
	}
}

func main() {

	var a, b int

	fmt.Print("Please input a:")
	fmt.Scanf("%d\n", &a)
	fmt.Print("\nPlease input b:")
	fmt.Scanf("%d\n", &b)
	fmt.Printf("\nSum from %d to %d is %d", a, b, sum(a, b))

	var c int
	fmt.Print("\nCount from 0 to c:")
	fmt.Scanf("%d\n", &c)
	fmt.Println()
	definiteLoop(&c)

	fmt.Println("Starting a infinite loop:")
	infiniteLoop()
}
