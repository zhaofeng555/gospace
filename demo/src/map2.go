package main

import (
	"fmt"
	"sort"
)

func test1() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}

	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)
}

func test2() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(m1)

	m2 := make(map[string]int)

	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}
func main() {
	test2()
}
