package main

import "fmt"

type Ainterface interface {
	Say()
}
type Binterface interface {
	Sang()
}

type child struct {

}

func (c child) Say(){
	fmt.Println("c中的say")
}

func (c child) Sang(){
	fmt.Println("c中的Sang")
}

type A interface {
	test01()
	test02()
}

type B interface {
	test01()
	test03()
}
type  C interface {
	A
	B
}
func main() {
	var C child

	var b Ainterface  = C

	b.Say()
	C.Sang()
	C.Say()

}
