// HelloWorld project main.go
package main

import (
	"fmt"
)

type myFloatFunc = func(float64, float64) float64

func myCalc(x float64) myFloatFunc {

	return func(y, z float64) float64 {
		return x + y + z
	}
}

func adder() func(int) int {
	sum := 0
	fmt.Println("sum:", sum)
	return func(x int) int {
		sum += x
		return sum
	}

}
func main() {
	fmt.Println(myCalc(10)(2, 3))

	pos := adder() // This is function closure
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}
	for i := 0; i < 10; i++ {
		fmt.Println(adder()(i))
	}
}
