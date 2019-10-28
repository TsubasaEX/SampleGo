package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprint("cannot Sqrt negative number:", float64(e))
	}
	return ""
}

func Sqrt(x ErrNegativeSqrt) (ErrNegativeSqrt, error) {
	return 0, x
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
