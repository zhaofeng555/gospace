package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Test()
	fmt.Println(a)

	b := B{}
	b.Test()
	fmt.Println(b)
}

func (a A) Test() {
	a.Name = "hhh"
	fmt.Println("A test()", a)
}

func (b *B) Test() {
	b.Name = "haojg"
	fmt.Println("B test()", b)
}
