package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func A(per *person) {
	per.Age = 13
	fmt.Println("A ", per)
}

type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
}
type student struct {
	human
	Name string
	Age  int
}

func main() {
	// a := person{
	// 	Name: "joe",
	// 	Age:  20,
	// }
	// fmt.Println(a)
	// a.Name = "haojg"
	// a.Age = 30
	// fmt.Println(a)
	// A(&a)
	// fmt.Println(a)

	a := teacher{Name: "joe", Age: 19, human: human{Sex: 0}}
	b := student{Name: "joe", Age: 20, human: human{Sex: 1}}

	fmt.Println(a, b)
	a.Age = 13
	a.Sex = 3
	fmt.Println(a, b)
}
