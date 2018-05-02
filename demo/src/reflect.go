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

func (u User) Hello(name string) {
	fmt.Println("hello ", name, ", my name is ", u.Name)

}

type Manager struct {
	User
	title string
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("no struct")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
}
func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("can't set")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("bad")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("byebye")
	}

}

func Set2() {
	u := User{1, "ok", 20}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}
func main() {
	// u := User{1, "ok", 20}
	// Info(u)

	// m := Manager{User: User{1, "OK", 13}, title: "me"}
	// t := reflect.TypeOf(m)

	// fmt.Printf("%v\n", t.FieldByIndex([]int{0, 1}))

	// u := User{1, "haha", 15}
	// Set(&u)
	// fmt.Println(u)

	Set2()

}
