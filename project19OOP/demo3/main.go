package main

import "fmt"

type Point struct {
	x string
	y int
}

type Rect struct {
	leftUp,rightUp Point
}
func main() {

	r1 := Rect{Point{"1",1},Point{"2",2}}
	fmt.Printf("l.x=%p l.y=%p  r.x=%p r.y=%p",&r1.leftUp.x,&r1.leftUp.y,&r1.rightUp.x,&r1.rightUp.y)
}

func test1(){
	for true{
		fmt.Println("test1")
	}
}
func test2(){
	for true{
		fmt.Println("test2")
	}
}
