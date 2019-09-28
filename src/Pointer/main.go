// SimpleWebApp project main.go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 15
	var b int = 25
	var p *int = &a
	p = &b

	fmt.Println(reflect.TypeOf(p))
	fmt.Println(p)
	fmt.Println(*p)
}
