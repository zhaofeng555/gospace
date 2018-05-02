package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	test7()
}

func test0() {
	a := person{}

	a.Name = "haojg"
	a.Age = 28
	fmt.Println(a)
	test1(a)
	fmt.Println(a)

	test2(&a)
	fmt.Println(a)
}

func test1(per person) {
	per.Age = 13
	per.Name = "luoli"
	fmt.Println("test1", per)
}

func test2(per *person) {
	per.Age = 33
	per.Name = "quejie"
	fmt.Println("test2", per)
}

func test3() {
	a := &person{
		Name: "hjg",
		Age:  20,
	}
	a.Name = "ok"
	fmt.Println(a)

	test2(a)
	fmt.Println(a)

}

//匿名结构
func test4() {
	a := &struct {
		Name string
		Age  int
	}{
		Name: "hhh",
		Age:  21,
	}
	fmt.Println(a)
}

//匿名结构
type user struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

func test5() {
	a := user{Name: "hjg", Age: 22}
	a.Contact.Phone = "1235"
	a.Contact.City = "beijing"
	fmt.Println(a)

	b := user{Name: "hjg", Age: 22}
	fmt.Println(a == b)
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

func test6() {
	a := teacher{Name: "t1", Age: 20, human: human{Sex: 0}}
	b := student{Name: "s1", Age: 12, human: human{Sex: 1}}
	fmt.Println(a)
	a.Age = 32
	a.Name = "laoshi"
	a.human.Sex = 100
	a.Sex = 90
	fmt.Println(a)

	fmt.Println(b)
}

type TZ int

func (tz *TZ) Increase(num int) {
	*tz += TZ(num)
}

func test7() {
	var a TZ
	a.Increase(100)
	fmt.Println(a)
}
