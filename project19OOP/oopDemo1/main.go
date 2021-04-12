package main

import "fmt"

func main()  {
	type structOne struct {
		name string
		age int
		sex string
		isStudent bool
	}
	a := structOne{
		 "lq",
		23,
		"man",
		 false,
	}
	//定义猫类
	type Cat struct {
		name string
		age int
		color string
		happy map[string]string
	}

	fmt.Println(a.sex)
	fmt.Printf("%T",a.sex)

	var cat1 Cat
	cat1.name = "小花"
	cat1.age = 3
	cat1.age = cat1.age >> 1
	fmt.Println(cat1)

	 monsterOnew := monster{
		 "牛魔王",
		 200,
		 []int{1},
	}

	monsterTwo := monsterOnew
	monsterTwo.name = "猪八戒"
	monsterTwo.slice = []int{2}
	fmt.Println(monsterOnew)
	fmt.Println(monsterTwo)

	fmt.Printf("%p",monsterTwo.slice)
	fmt.Printf("%p",monsterOnew.slice)

}
type monster struct {
	name string
	age int
	slice []int
}