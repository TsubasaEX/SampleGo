// HelloWorld project main.go
package main

import (
	"fmt"
	"math"
)

type myFloat float64
type Vertex struct {
	x, y int
}

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f int) {
	v.x *= f
	v.y *= f
}
func main() {
	vf := myFloat(-math.Sqrt2)
	fmt.Println(vf.Abs())

	p := &Vertex{3, 4}
	p.Scale(10)
	fmt.Println(p.x*p.x, p.y*p.y)
}
