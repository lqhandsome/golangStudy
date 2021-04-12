package main

import (
	"fmt"
)

type SourceMonkey struct {
	name string
}

type Monkey struct {
	SourceMonkey
}

func (mon *Monkey)Use()  {
	fmt.Println("这个猴子叫=",mon.name)
}

func (mon *Monkey)Swimming()  {
	fmt.Println("这个猴子可以游泳",mon.name)
}


type Fish interface {
	Swimming()
}

type Bird interface {
	Flying()
}
type interfaceArr [6]interface{}
func main() {

	var a interfaceArr
	for i,_ := range a{
		a[i] = i + 1
	}
	fmt.Println(a)
	monkey := Monkey{
		SourceMonkey{
			"悟空",
		},
	}
	monkey.Use()
	monkey.Swimming()
	go say("world")
	say("hello")

}


func say(say string){
	for i := 0;i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(say)
	}
}