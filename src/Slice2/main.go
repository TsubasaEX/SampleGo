// Slice2 project main.go
package main

import (
	"fmt"
)

func main() {

	a := make([]int, 5)
	b := make([]float64, 6, 6)
	c := b[2:]
	fmt.Printf("%d %d %v\n", len(a), cap(a), a)
	fmt.Printf("%d %d %v\n", len(b), cap(b), b)
	fmt.Printf("%d %d %v\n", len(c), cap(c), c)
	fmt.Println("Hello World!")

	s := [][]int{
		[]int{1, 2, 3},
		append(make([]int, 5), 4, 5, 6)} // slices of slices
	fmt.Println(len(s), cap(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(len(s[i]), cap(s[i]), s[i])
	}
	for i, v := range s {
		fmt.Println(i, ":", len(v), cap(v), v)
	}
	ary := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	for i, v := range ary {
		fmt.Println(i, v)
	}
	fmt.Println(ary)
}
