package main

import (
	"fmt"
)

func test1() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a = 0")
	case 1:
		fmt.Println("a = 1")
	default:
		fmt.Println("none")
	}
	switch {
	case a >= 0:
		fmt.Println("a >= 0")
		fallthrough
	case a >= 1:
		fmt.Println("a >= 1")
	default:
		fmt.Println("none")
	}
}

func test2() {
LABLE1:
	for {
		for i := 0; i < 10; i++ {

			if i > 3 {
				break LABLE1
			}
		}
	}
	fmt.Println("ok")
}
func test3() {
	c := [...]int{99: 1}
	var p *[100]int = &c
	fmt.Println(p)
}
func test4() {
	x, y := 1, 2
	a := [...]*int{&x, &y}
	fmt.Println(a)
}
func test5() {
	a := [10]int{}
	a[1] = 2
	fmt.Println(a)
	p := new([10]int)
	p[1] = 3
	fmt.Println(p)
}
func test6() {
	a := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Println(a)
}
func main() {
	test6()

}
