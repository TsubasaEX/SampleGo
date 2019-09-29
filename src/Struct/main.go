// Struct project main.go
package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
	id   string
}

type car struct {
	gas_petal   int
	brake_petal uint16
	steer       int16
	speed       float64
	man         person
}

func main() {
	var a_car = car{97, 20, -100, 12.256, person{"andy", "1"}}
	b_car := car{gas_petal: 100,
		brake_petal: 200,
		steer:       -1000,
		speed:       225.25}

	pb := &b_car
	var a_str string = strconv.Itoa(a_car.gas_petal)

	fmt.Println(a_car)
	fmt.Println(a_str)
	fmt.Println(a_car.man.name)
	fmt.Println(b_car.speed)
	pb.speed = 300
	fmt.Println(b_car.speed)
	fmt.Println(pb.man.name)

	f, err := strconv.ParseFloat("2.256", 64)
	g, err2 := strconv.ParseInt("16", 10, 32)
	h, err3 := strconv.ParseInt("10000", 2, 32)
	fmt.Println(f, err)
	fmt.Println(g, err2)
	fmt.Println(h, err3)
}
