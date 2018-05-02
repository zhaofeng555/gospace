package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	test5()
}
func test5() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("time out")
	}
}
func test4() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)

			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1)
	close(c2)

	for i := 0; i < 2; i++ {
		<-o
	}
}
func test3() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go2(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}
func Go2(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true
}
func test2() {
	c := make(chan bool)
	go func() {
		fmt.Println("go go go!")
		c <- true
	}()
	fmt.Println("go 1")
	<-c
	fmt.Println("go 2")
}
func test1() {
	c := make(chan bool)
	go func() {
		fmt.Println("let us go!")
		c <- true
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
func test0() {
	go Go()
	time.Sleep(2 * time.Second)
}
func Go() {
	fmt.Println("Go Go Go !!!")
}
