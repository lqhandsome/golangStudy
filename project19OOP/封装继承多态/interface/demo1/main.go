package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}
type Phone struct {

}
func (p Phone)Start(){
	fmt.Println("Phone工作了")
}
func (p Phone)Stop(){
	fmt.Println("Phone停止工作了")
}
type Camera struct {

}
type Computer struct {

}
func (c Computer) Working(usb Usb){
	usb.Start()
	usb.Stop()
}
func (c Camera)Start(){
	fmt.Println("相机工作了")
}
func (c Camera)Stop(){
	fmt.Println("相机停止工作了")
}
func main() {
	p := Phone{}
	c := Camera{}
	cum := Computer{}
	cum.Working(p)
	cum.Working(c)
}
