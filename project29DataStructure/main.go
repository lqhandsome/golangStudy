package main

import (
	"fmt"
	"strings"
)

type Foo struct {
	Name string
}

var bar = "hello biezhi"

// -------------方法返回值----------------
func returnValue() string {
	return bar
}

func returnPoint() *string {
	return &bar
}

// --------------方法入参-----------------
func modifyNameByPoint(foo *Foo) {
	foo.Name = "emmmm " + foo.Name
}

func nameToUpper(foo Foo) string {
	foo.Name = strings.ToUpper(foo.Name)
	return foo.Name
}

// --------------实例方法-----------------
func (foo Foo) setName(name string) {
	foo.Name = name
}

func (foo *Foo) setNameByPoint(name string) {
	foo = &Foo{}
}

func main() {
	fooPoint := &Foo{Name: "lq"}
	modifyNameByPoint(fooPoint)
	fmt.Println(fooPoint) //&{emmmm lq} 会修改到原值

	foo := Foo{Name: "hhp"}
	//nameToUpper(foo)
	//fmt.Println(foo) //{hhp}
	//fmt.Println(nameToUpper(foo)) //HHP

	foo.setName("aaa")
	fmt.Println(foo)

	foo.setNameByPoint("55")
	fmt.Println(foo)


	fooPoint2 := &Foo{Name: "lq1"}
	fooPoint2.setNameByPoint("aaa")
	fmt.Println(fooPoint2)
	fooPoint2.setName("nbb")
	fmt.Println(fooPoint2)
}
