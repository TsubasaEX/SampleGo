// Method project main.go
package main

import (
	"fmt"
	"strings"
)

const ratio = 2.5

type weapon struct {
	attack  float64
	defense float64
}

type person struct {
	attack  float64
	defense float64
	power   float64
	arm     *weapon
}

func (p person) getPowerRatio() float64 { //Value Receiver
	return p.power * ratio
}

func (p person) updateAttack(a float64) { //Value Receiver can't modify struct values
	p.attack = a
	p.power = p.attack * p.defense
}

func (p *person) updateDefense(d float64) {
	p.defense = d
	p.power = p.attack * p.defense
}

func main() {
	man := person{attack: 10.1,
		defense: 20.2,
		power:   10.1 * 20.2,
		arm:     &weapon{10, 20}}
	fmt.Println(man)
	fmt.Println(man.getPowerRatio())
	man.updateAttack(20.2)
	fmt.Println(man.getPowerRatio())
	man.updateDefense(40.4)
	fmt.Println(man.getPowerRatio())

	pm := &man
	var a int = 1
	pa := &a
	fmt.Println(*pa)
	fmt.Println(pm.attack)
	fmt.Println(pm.arm.attack)

	ia := [10]int{}
	for i := 0; i < len(ia); i++ {
		ia[i] = i
	}
	for i := 0; i < len(ia); i++ {
		fmt.Println(ia[i])
	}
	for key, value := range ia {
		fmt.Println(key)
		fmt.Println(value)
	}
	fighters := [3]person{{10.1, 20.2, 10.1 * 20.2, &weapon{10, 20}},
		{10.1, 20.2, 10.1 * 20.2, &weapon{11, 20}},
		{10.1, 20.2, 10.1 * 20.2, &weapon{12, 20}}}
	enmies := []person{{10.1, 20.2, 10.1 * 20.2, &weapon{10, 20}}, {10.1, 20.2, 10.1 * 20.2, &weapon{11, 20}}}
	for key, value := range fighters {
		fmt.Println(key)
		fmt.Println(value)
		fmt.Println(value.arm.attack)
	}
	for key := range fighters {
		fmt.Println(key)
	}
	fmt.Println(fighters)
	for _, value := range enmies {
		fmt.Println(value)
	}
	fmt.Println(enmies)

	str := []string{"abcde", "fghij"}
	pstr := &str
	fmt.Printf("%c\n", str[1][1])
	fmt.Printf("%T\n", pstr)
	fmt.Println((*pstr)[1][1])
	(*pstr)[1] = strings.Replace((*pstr)[1], "g", "z", -1)
	// (*pstr)[1][1] = "z" // not working (immutable)
	fmt.Println(str[1])

	str2 := "a space-separated string"
	str2 = strings.Replace(str2, " ", ",", -1)
	fmt.Println(str2)
}
