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

// Type Assertion
func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func assert2(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var inf PayInterface
	b := Boss{100, 0.2, "A"}
	w := Worker{10, "B"}

	fmt.Println(b.getPay(10))
	fmt.Println(w.getPay(10))

	inf = b
	fmt.Println(inf.getPay(10), inf.(Boss), inf.(Boss).salary)
	fmt.Printf("Interface type %T value %v\n", inf, inf)
	inf = w
	fmt.Println(inf.getPay(10), inf.(Worker), inf.(Worker).salary)
	fmt.Printf("Interface type %T value %v\n", inf, inf)

	// inf = &b
	// fmt.Println(inf.getPPay())
	// inf = &w
	// fmt.Println(inf.getPPay())

	var s interface{} = 56
	assert(s)
	assert2(s)
	var ss interface{} = "Steven"
	assert(ss)
	assert2(ss)
}
