package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string `json:"monster_name"`
	Age int
}
func main() {
	monster := Monster{
		"{\"Name\":\"牛魔王\",\"Age\":34}",
		34,
	}

	jsonStruct, err := json.Marshal(monster)

	if err != nil {
		fmt.Println("转换失败")
	}

	fmt.Println(string(jsonStruct))
}
