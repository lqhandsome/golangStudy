package main

import (
	 "encoding/json"
	"fmt"
)
func main() {
	//结构体之间转化,必须字段类型相同,名称相同,顺序相同,个数相同
	var a A
	var b B
	a = A(b)
	fmt.Println(a,b)
	monster := Monster{
		"红孩儿",
		300,
		"三味真火",
	}
	json,_ := json.Marshal(monster)
	fmt.Println(string(json))
}

type Monster struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	Skill string `json:"skill"`
}
type A struct {
	age int
	name string
}

type B struct {
	age int
	name string
}