package main

import "fmt"

func main() {
	m := aa()
	fmt.Printf("%T",m)
}

type Monster struct {
	UserId int `json:"userId"`
	UserName string `json:"userName"`
	UserPwd string `json:"userPwd"`

}

func aa() (m *Monster){
	m = new(Monster)
	print(m)
	fmt.Println()
	fmt.Printf("m =%p\n",m)
	m.UserId = 1
	return
}

func bb() (m map[string]string) {

	return m
}