// HelloWorld project main.go
package main

import (
	"fmt"
)

type PayInterface interface {
	getPay(float64) float64
	// getPPay() float64	//can't use value receiver and pointer receiver at same time
}

type Boss struct {
	salary   int
	ratio    float64
	division string
}

type Worker struct {
	salary   int
	division string
}

func (b Boss) getPay(f float64) float64 {
	return float64(b.salary) * b.ratio * f
}

// func (b *Boss) getPPay() float64 {
// 	return float64(b.salary) * b.ratio
// }

func (w Worker) getPay(f float64) float64 {
	return float64(w.salary) * f
}

// func (w *Worker) getPPay() float64 {
// 	return float64(w.salary)
// }

func main() {
	var inf PayInterface
	b := Boss{100, 0.2, "A"}
	w := Worker{10, "B"}

	fmt.Println(b.getPay(10))
	fmt.Println(w.getPay(10))

	inf = b
	fmt.Println(inf.getPay(10))
	inf = w
	fmt.Println(inf.getPay(10))

	// inf = &b
	// fmt.Println(inf.getPPay())
	// inf = &w
	// fmt.Println(inf.getPPay())
}
