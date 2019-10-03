// Slice project main.go
package main

import (
	"fmt"
	"reflect"
)

func main() {

	ary := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := ary[1:3]
	fmt.Println(reflect.TypeOf(ary)) //array
	fmt.Println(reflect.TypeOf(s))   //slice
	fmt.Println(ary)
	fmt.Println(s)
	s[1] = 30
	fmt.Println(ary)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Printf("%p %p\n", &ary, s)
	fmt.Println(&ary, &ary[0])
	fmt.Println(&ary[1], &s[0])
	s2 := ary[1:]
	fmt.Println(s2)
	s3 := ary[:3]
	fmt.Println(s3)
	s4 := ary[:]
	fmt.Println(reflect.TypeOf(s4))

	man := struct {
		attack  int
		defense int
	}{10, 20}
	fmt.Println(man)
	fmt.Println(reflect.TypeOf(man))

	people := []struct {
		age     int
		finance int
	}{{21, 2100}, {35, 3500}}
	fmt.Println(people)

	slice2 := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Println(slice2)
	sa := slice2[:0]
	fmt.Printf("len : %d, cap: %d, value: %v\n", len(sa), cap(sa), sa)
	sa = slice2[:3]
	fmt.Printf("len : %d, cap: %d, value: %v\n", len(sa), cap(sa), sa)
	sa = slice2[1:]
	fmt.Printf("len : %d, cap: %d, value: %v\n", len(sa), cap(sa), sa)
	fmt.Println("After Insert 'G'......")
	slice2 = append(slice2, "G") //this will increase the cap!! Because the cap is not enough
	slice2 = append(slice2, "H") // this won't increase the cap.
	fmt.Println(slice2)
	sa = slice2[:]
	fmt.Printf("len : %d, cap: %d, value: %v\n", len(sa), cap(sa), sa)

	slice3 := make([]string, 2, 2)
	sa2 := slice2[1:4]
	fmt.Println(slice3)
	fmt.Println(sa2)
	copy(slice3, sa2)
	fmt.Println(slice3)
	slice3[0] = "BB"
	fmt.Println(slice2)
	sa2[0] = "BB"
	fmt.Println(slice2)
	fmt.Println("Hello World!")
}
