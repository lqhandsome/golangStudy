package main

import "fmt"

type Student struct {
	name string
	age int
}
type Popil struct {
	Student
	Age int
}
func (stu Student)print(){
	fmt.Println("stu")
}
func (stu Popil)print(){
	fmt.Println("Popil")
}


type Tv struct {
	Goods
	Brand
}
type Tv2 struct {
	*Goods
	*Brand
}

type Goods struct {
	name string
	price float64
}

type Brand struct {
	name string
	Address string
}
func main() {
	//s := Popil{}
	//s.age = 10
	//s.Student.print()
	//fmt.Println(s)
	tv1 := Tv{
		Goods {
			"电视机",
			100,
		},
		Brand {
			"小米",
			"北京",
		},
	}
	tv2 := Tv2{
		&Goods {
			"电视机",
			100,
		},
		&Brand {
			"小米",
			"北京",
		},
	}
	tv11 := tv1
	tv22 := tv2
	tv1.Goods.name ="aaa"
	tv2.Goods.name ="bbb"
	fmt.Println("tv1=",tv1)
	fmt.Println("tv11=",tv11)
	fmt.Println(*tv2.Goods)
	fmt.Println(*tv22.Goods)
}
