package main

import "fmt"

func main()  {
	 stu := make(map[string]map[string]string,10)
	 var name string
	 fmt.Scanln( &name)
	stu["lq"] = make(map[string]string,2)
	stu["lq"]["pwd"] ="666666"
	stu["lq"]["nickName"] ="nickName"
	 modify(stu,name)
	 fmt.Println(stu)
}
func modify(stu map[string]map[string]string,name string)  {
	if stu[name] != nil {
	//有这个用户
		stu[name]["pwd"] = "888"
	} else {
		stu[name] = make(map[string]string,2)
		stu[name]["pwd"] ="666666"
		stu[name]["nickName"] ="nickName"
	}
}