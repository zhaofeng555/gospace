package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("connected:", pc.name)
}

type TVConnecter struct {
	name string
}

func (tc TVConnecter) Name() string {
	return tc.name
}

func (tc TVConnecter) Connect() {
	fmt.Println("connected:", tc.name)
}

func Disconnect(usb interface{}) {

	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("disconneted PhoneConnecter :", v.name)
	case TVConnecter:
		fmt.Println("disconneted TVConnecter :", v.name)
	default:
		fmt.Println("haojg device")
	}

	// if pc, ok := usb.(PhoneConnecter); ok {
	// 	fmt.Println("Disconnected:", pc.name)
	// 	return
	// }

}

func main() {
	pc := TVConnecter{"TVConnecter"}
	var a Connecter
	a = USB(pc)
	a.Connect()
	Disconnect(a)

	var b interface{}
	fmt.Println(b == nil)

	var p *int = nil
	b = p
	fmt.Println(b == nil)
}
