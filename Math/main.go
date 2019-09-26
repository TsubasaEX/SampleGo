// Math project main.go
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func calc(a float64, b float64) float64 {
	var pow = math.Pow(a, b)
	fmt.Printf("The power(%f,%f): %f\n", a, b, math.Pow(a, b))
	fmt.Println("The square of 4:", math.Sqrt(4))
	fmt.Printf("The rand number in 1-100: %d\n", rand.Intn(100)+1)
	return pow * 2
}

func main() {
	var str1, str2 = "str1", "str2"
	var str3 string
	str3 = "str3"
	var seed = time.Now().UnixNano()
	fmt.Println(str1, str2, str3)
	fmt.Println(time.Now())
	fmt.Println("seed:", seed)
	rand.Seed(seed)
	var pow2 = calc(2, 4)
	fmt.Println(pow2)
}
