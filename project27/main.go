package main

import "fmt"

func main() {
	//monster := Monster{
	//	UserId:   0,
	//	UserName: "",
	//	UserPwd:  "",
	//	Animal:   Animal{},
	//}
	//
	var monster Monster
	monster.UserId = 1
	monster.UserPwd = "123445"
	monster.Animal.Name = "lq"

	fmt.Println(monster)
	monster.Animal.pp()
}

type Monster struct {
	UserId  int    `json:"userId"`
	UserPwd string `json:"userPwd"`
	Animal  Animal
}

type Animal struct {
	Name string
}

func (a Animal) pp() {
	fmt.Println(a.Name)
}
