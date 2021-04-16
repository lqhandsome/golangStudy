package main

import (
	"fmt"
	"encoding/json"
)
type Monster struct {
	//monster_name string
	Name string
	Age int


}

func main() {
	//decodeJsonStruct()
	decodeJsonSlice()
}

func decodeJsonStruct() {
	var monster Monster
	jsonStr := "{\"Name\":\"牛魔王\",\"Age\":34}";
	err := json.Unmarshal([]byte(jsonStr),&monster)

	if err != nil {
		fmt.Println("解析失败")
	}
	fmt.Println(monster)
}

func decodeJsonSlice() {
		var s []map[string]interface{}

	jsonStr := "[{\"Name\":\"牛魔王\",\"Age\":34},{\"Name\":\"牛魔王\",\"Age\":34}]";
	err := json.Unmarshal([]byte(jsonStr),&s)
	if err != nil {
		fmt.Println("解析失败")
	}
	fmt.Println(s)
}