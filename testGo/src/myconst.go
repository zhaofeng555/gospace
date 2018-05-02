package main

import (
	"fmt"
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
)

func main() {
	a := 10
	a++
	var p *int = &a
	fmt.Println(*p)

	//b := 2
	for i := 0; i < 5; i++ {
		i++
		fmt.Println(i)
	}
	fmt.Println("over")
}
