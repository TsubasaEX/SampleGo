// HelloWorld project main.go
package main

import (
	"fmt"
)

func sum(nums ...int) int {
	fmt.Println(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
}
