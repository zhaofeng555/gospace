package main

import (
	"fmt"
)

func main() {
	var m = sum([]int{1,2,3})
	fmt.Println(m)
}
func sum(a []int) int{
	s:=0
	for i := 0; i < len(a); i++{
		s += a[i]
	}
	return s;
}

