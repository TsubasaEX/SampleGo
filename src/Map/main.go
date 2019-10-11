// Map project main.go
package main

import (
	"fmt"
)

type myMap struct {
	lat, long string
}

func main() {

	var m map[string]myMap
	m = make(map[string]myMap)
	m["John"] = myMap{"40", "40"}

	m2 := map[string]myMap{"Amy": {"10", "10"}, "Alex": {"20", "20"}}
	m3 := map[string]myMap{"Kyle": myMap{"30", "30"}}
	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)

	m2["Amy"] = myMap{"11", "11"}
	fmt.Println(m2)
	fmt.Println(m2["Amy"])

	m2["Zach"] = myMap{"50", "50"}
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	// fmt.Println(len(m2), m2)

	delete(m2, "Alex")
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	v, ok := m2["Zach"]
	fmt.Println("The value:", v, "Present?", ok)
	v, ok = m2["Alex"]
	fmt.Println("The value:", v, "Present?", ok)
	// fmt.Println(len(m2), m2)

	normap := make(map[string]int)
	normap["A"] = 1
	normap["B"] = 2
	fmt.Println(normap)

	normap2 := map[string]int{"C": 3, "D": 3}
	fmt.Println(normap2)
}
