package main

import "fmt"

func main()  {
	var nil_map map[string]string
	println(nil_map)

	emp_map := map[string]string{}

	 student := make(map[int]map[string]string)
	println(emp_map)
	println(student)
	student[1] = make(map[string]string)
	student[1]["name"] = "lq"
	student[1]["sex"] = "man"
	student[2] = make(map[string]string)
	student[2]["name"] = "hhp"
	student[2]["sex"] = "woman"
	findValue,ok := student[1]["namae"]
	if findValue == ""{
		fmt.Println("not have")
	}
	 fmt.Println(findValue,ok)
}
