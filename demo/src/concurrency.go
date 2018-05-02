package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// testGo()
	// testGo2()
	test4()
}

var c chan string

func test4() {
	c = make(chan string)
	go Pingpong()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("from main : hello, #%d", i)
		fmt.Println(<-c)
	}
}
func Pingpong() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("from pingpong: hi, #%d", i)
		i++
	}
}
func test3() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("time out ")

	}
}

func testGo() {
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}

func testGo2() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go2(&wg, i)
	}
	wg.Wait()
}
func Go2(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}
