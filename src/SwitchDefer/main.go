// Switch project main.go
package main

import (
	"fmt"
	"myMathSecond"
	"myMathSecond/myLib/myMath"
	"runtime"
)

func Sum(a int, b int) {
	var sum int

	defer fmt.Println("defer Sum:", sum)
	sum = a + b
	fmt.Println("Sum:", sum)

}

func main() {

	defer Sum(10, 20)
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println(os)
	}

	a := "Z"
	t := "Zach"

	switch t {
	case a + "ach":
		fmt.Println("Got Zach")
	default:
		fmt.Println("Default")
	}

	switch {
	case a == t:
		fmt.Println("a == t")
	case a+"T" == t:
		fmt.Println("a+\"T\"==t")
	case a+"ach" == t:
		fmt.Println("a+\"ach\"==t")
	default:
		fmt.Println("Default")
	}

	mul := myMathSecond.Multiply_Second(100, 1.11)
	fmt.Println(mul)

	sum := myMath.MySum(1.1, 2.2)
	fmt.Println(sum)
}
