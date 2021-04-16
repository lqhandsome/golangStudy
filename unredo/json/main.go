package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

func jsonStruct(){
	monster := Monster{
		"牛魔王",
		500,
		"芭蕉扇",
	}

	jsonMonster, err := json.Marshal(monster)

	if err != nil {
		fmt.Println("json错误")
	}
	fmt.Println(monster)
	fmt.Printf("%v",string(jsonMonster))
}

func jsonMap() {
	var m map[string]interface{}
	print(m)
	m  = make(map[string]interface{})
	print("\n")
	print(m)
	m["name"] = "lq"
	m["age"] = 23
	m["address"] = [3]string{"a","b","c"}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("转化失败")
	}
	fmt.Println("jsonMap",string(data))
}

func jsonSlice() {
	var a []map[string]interface{}
	//a = make([]map[string]interface{},5)
		print(a)
	m1 := make(map[string]interface{})
	m1["name"] = "lq"
	m1["age"] = 23
	m1["address"] = [3]string{"a","b","c"}

	a = append(a,m1)

	m2 := make(map[string]interface{})
	m2["name"] = "lqqq"
	m2["age"] = 100
	m2["address"] = [3]string{"aa","bb","cc"}
	a = append(a,m2)

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("转化失败")
	}

	fmt.Println("jsonMap",string(data))
	fmt.Println(len(a),cap(a))
	print(a)

}

func  main()  {
	jsonStruct()
	jsonMap()
	jsonSlice()
}