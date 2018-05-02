package main

import (
	"fmt"
)

func main() {
	test3()
}
func test5(){
	s1:=make([]int, 3, 6)
	fmt.Println("%p\n", s1)
	s1=append(s1, 1, 2, 3)
	fmt.Println("%v %p\n", s1, s1)
	s1=append(s1, 1, 2, 3)
	fmt.Println("%v %p\n", s1, s1)
}
func test4() {
	s2 := make([]int, 3, 30)
	fmt.Println(s2)

	fmt.Println("hello", "world")
	sum, prod := learnMultiple(2, 3)
	fmt.Println(sum, prod)

	sum1 := test2(4)
	fmt.Println(sum1)
}
func test3() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
	sa := a[2:5]
	fmt.Println(sa)
	fmt.Println(len(sa), cap(sa))
	sa2:=a[:]
	fmt.Println(sa2)
}
func test1() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	s1 := a[5:8]
	fmt.Println(s1)
	fmt.Println(a[:6])
	fmt.Println(a[5:])
	fmt.Println(a[3:6])
}
func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}
func test2(x int) (sum int) {
	return x * x
}
