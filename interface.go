package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Connector
}

type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("connect : ", pc.name)
}

func main() {
	var a USB
	a = PhoneConnector{"PhoneConnector"}
	a.Connect()
	Disconnect(a)
}

func Disconnect2(usb USB) {
	if pc, ok := usb.(PhoneConnector); ok {
		fmt.Println("Disconnected:", pc.name)
		return
	}
	fmt.Println("Unknown decive.")
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("unknown decive")
	}
}
