// HelloWorld project main.go
package main

import (
	"fmt"
	"math"
)

type myFloatFunc = func(float64, float64) float64
type myFloat = float64

func compute(a float64, b float64, fn func(float64, float64) float64) float64 {
	return fn(a, b)
}

func compute2(a myFloat, b myFloat, fn myFloatFunc) myFloat {
	return fn(a, b)
}

func mySqrt(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {
	s := compute(3, 4, mySqrt)
	fmt.Println(s)

	lf := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(3, 4, lf))

	fmt.Println(compute(3, 4, math.Pow))

	fmt.Println(compute2(3, 4, math.Pow))
}
