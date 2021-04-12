package main

import "fmt"

func main() {
	phone := Phone{
		"新款iphone",
		69999,
	}

	P(phone)
}
	type Usb interface {
		Start()
		Stop()
	}

	type Phone struct {
		name string
		money float64
	}

	func (p Phone) Start(){
		fmt.Println("手机开始工作")
	}
	func (p Phone) Stop(){
		fmt.Println("手机停止工作")
	}

	func P(u Usb){
		u.Start()
		u.Stop()
	}