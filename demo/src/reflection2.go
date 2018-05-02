package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func test0() {
	m := Manager{User: User{1, "ok", 12}, title: "123"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
}
func test2() {
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)

	fmt.Println(x)
}
func test3() {
	u := User{1, "ok", 12}
	fmt.Println(u)
	Set(&u)
	fmt.Println(u)
}
func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("bad")
		return
	} else {
		fmt.Println("good")
	}

	if f.Kind() == reflect.String {
		f.SetString("bye bye")
	}
}

func (u User) Hello(name string) {
	fmt.Println("hello ", name, ", my name is", u.Name)
}
func test4() {
	u := User{1, "hjg", 22}
	fmt.Println(u)

	u.Hello("gg")

	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("jack")}
	mv.Call(args)
}
func main() {
	test4()
}
