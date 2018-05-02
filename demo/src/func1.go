package main

import (
	"fmt"
)

func main() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func A(a int, b string) int {
	return 1
}

func B(a, b, c int) (d, e int) {
	d, e = 1, 2
	return
}
func D(a, b, c int) (int, int) {
	d, e := 1, 2
	return d, e
}
func E() int {
	return 1
}

//a slice 形参，值拷贝
func F(a ...int) {
	fmt.Println(a)
}

//s 会修改原来的数组（地址拷贝）
func G(s []int) {
	s[0] = 9
	s[1] = 8
}

func H() {
	//匿名函数
	a := func() {
		fmt.Println("Func A")
	}
	a()
}

//闭包
func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func test3() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	/*
		结果：
		a
		c
		b
	*/
}

func test4() {
	//注册recover函数
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()

	//程序终止
	panic("Panic in B")
}
