package main

import "fmt"

func main()  {
	//map切片
	var mapOne []map[string]string
	mapOne = make([]map[string]string,5)
	mapOne[0] = make(map[string]string)
	mapOne[0]["name"] = "lq"

	fmt.Println(mapOne)
	updateMap(mapOne[0])
	fmt.Println(mapOne)
	type Stu struct {
		name string
		age int
		sex string
	}
}

func updateMap(mm map[string]string)  {
		mm["name"] = "newName"
}
