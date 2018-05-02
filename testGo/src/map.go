package main

import (
	"fmt"
)

func main() {
	test6();
}

func test1() {
	var m map[int]string
	m = make(map[int]string)
	fmt.Println(m)
}
func test2() {
	var m map[int]string
	m = map[int]string{}
	fmt.Println(m)
}
func test3() {
	var m map[int]string=make(map[int]string)
	fmt.Println(m)
}
func test4() {
	m:=make(map[int]string)
	m[1]="hello"
	m[2]="world"
	fmt.Println(m[1])
	fmt.Println(m[3])
	fmt.Println(m)
	delete(m, 1);
	fmt.Println("m[1] = ",m[1])
}
func test5(){
	var m map[int]map[int]string
	m=make(map[int]map[int]string)
	a, ok := m[2][1]
	fmt.Println(a, ok)
	if !ok {
		m[2]=make(map[int]string)
	}
	m[2][1]="Good"
	a, ok = m[2][1]
	fmt.Println(a, ok)
}
func test6(){
	sm:=make([]map[int]string, 5)
	for _, v := range sm{
		v =make(map[int]string, 1)
		v[1]="ok"
		fmt.Println(v);
	}
	fmt.Println(sm);
	
	sm2:=make([]map[int]string, 5)
	for i := range sm2{
		sm2[i] =make(map[int]string, 1)
		sm2[i][1]="ok"
		fmt.Println(sm2[i]);
	}
	fmt.Println(sm2);
}