// Type project main.go
package main

import (
	"fmt"
	"reflect"
)

const c1, c2 int = 10, 20

const (
	s1 string = "str1"
	s2 string = "str2"
)

var (
	i1 = 1
	i2 = 2
)

var io1, io2 float64 = 4, 5

func swap(x, y string) (string, string) {
	return y, x
}

func swap2(x, y string) (b, a string) {
	a, b = x, y
	return
}

func main() {

	var i3 = 3

	fmt.Println("Type:", reflect.TypeOf(c1), "Value:", c1)
	fmt.Printf("Type: %T Value: %f\n", float64(c2), float64(c2))
	fmt.Println(s1, s2)
	fmt.Println(i1, i2)
	fmt.Println(i3, io1, io2)
	fmt.Println(swap("abc", "def"))
	fmt.Println(swap2("abc", "def"))
}
