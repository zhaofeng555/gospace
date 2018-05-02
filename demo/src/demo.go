package main

import (
	"fmt"
	"sort"
	"unsafe"
)

func varFun() {
	var a int
	const (
		LEN  = 10
		WITH = 12
		str  = "abc"
		b    = iota
		c
	)
	a = 10
	fmt.Println(a)
	a = 11
	fmt.Println(a)
	fmt.Println(LEN * WITH)
	fmt.Println(str, len(str), unsafe.Sizeof(str))
	fmt.Println(b, c)
}
func optFun() {
	var a int = 12
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("a - b =  %d\n", c)
	c = a * b
	fmt.Printf("a*b = %d\n", c)

	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a <= b")
	}

	var d bool = true
	var e bool = false
	if d && e {
		fmt.Println("d && e = true")
	}
	if d || e {
		fmt.Println("d || e = true")
	}
}

func switchFun() {
	var grade string = "A"

	switch grade {
	case "A":
		fmt.Println("优")
	case "B":
		fmt.Println("良")
	default:
		fmt.Println("---")
	}

	switch {
	case grade == "A":
		fmt.Println("优")
	case grade == "B":
		fmt.Println("良")
	default:
		fmt.Println("---")
	}

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型 ：%T", i)
	case int:
		fmt.Printf("x 是 int")
	case string:
		fmt.Printf("x 是 string")
	default:
		fmt.Printf("未知")
	}
}
func swap(x, y string) (string, string) {
	return y, x
}
func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum / size)
	return avg
}
func sortFun() {
	a := [...]int{2, 6, 3, 8, 5, 4, 9}
	fmt.Println(a)

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}

func slicFun() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	s1 := a[5:]
	fmt.Println(s1)
}
func slicFun2() {
	s1 := make([]int, 3, 6)
	fmt.Printf("%p \n", s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
}
func slicFun3() {
	a := []int{1, 2, 3, 4, 5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2)
	s1[0] = 9
	fmt.Println(s1, s2)

	b := []int{6, 7, 8}
	copy(a, b)
	fmt.Println(a)
}
func mapFun() {
	var m map[int]string
	m = map[int]string{}
	fmt.Println(m)

	m = make(map[int]string)
	fmt.Println(m)

	m1 := make(map[int]string)
	fmt.Println(m1)
	m1[1] = "OK"
	fmt.Println(m1)
	m1[2] = "Hello"
	fmt.Println(m1)
	delete(m1, 2)
	fmt.Println(m1)

	var m3 map[int]map[int]string
	m3 = make(map[int]map[int]string)
	a, ok := m3[2][1]
	fmt.Println(a, ok)
	if !ok {
		m3[2] = make(map[int]string)
	}
	m3[2][1] = "good"
	a, ok = m3[2][1]
	fmt.Println(a, ok)

	m4 := make([]map[int]string, 5)
	for _, v := range m4 {
		v = make(map[int]string, 1)
		v[1] = "OK"
		fmt.Println(v)
	}
	fmt.Println(m4)

	m5 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m5))
	i := 0
	for k, _ := range m5 {
		s[i] = k
		i++
	}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)
}

func myFun(a int, b string) (c int) {
	fmt.Println(b)
	c = a + 1
	return c
}

func myFun2(a *int) {
	*a = 2
	fmt.Println(*a)
}

func myFunType() {
	fmt.Println("print function  type")

	f := func() {
		fmt.Println("inner func ")
	}
	f()
}

//闭包函数
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

func deferFun() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}
func main() {
	//varFun()
	// optFun()
	// switchFun()
	// a, b := swap("hello", "world")
	// fmt.Println(a, b)
	// var blance = []int{1000, 2, 3, 17, 50}
	// var avg float32
	// avg = getAverage(blance, 5)
	// fmt.Printf("平均值：%f", avg)
	// sortFun()
	// slicFun3()
	// mapFun()

	// c := myFun(2, "a")
	// fmt.Println(c)

	// a := 1
	// myFun2(&a)
	// fmt.Println(a)

	// t := myFunType
	// t()

	// f := closure(10)
	// fmt.Println(f(1))

	deferFun()
}
