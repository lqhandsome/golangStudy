package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	var s []int
	s = append(s,1)

	println(s)
	fmt.Println(s)
	return
	a := 1
	b := make([]int,2)
	b[1] = 10
	println(b)
	b = append(b,1)
	b = append(b,1)
	b = append(b,1)

	println(a)
	println(b)
	//c := [2]int{1,2}
	//println(c)
	return
	m := aa()
	m.UserId  = 100
	m.UserName  = "lqhandsome"
	m.UserPwd  = "123456"
	conn,_ := redis.Dial("tcp","127.0.0.1:6379")
	res,_ := json.Marshal(m)
	conn.Do("hset","user",m.UserId,string(res))
	fmt.Printf("%p%p",m,&m)
}

type Monster struct {
	UserId int `json:"userId"`
	UserName string `json:"userName"`
	UserPwd string `json:"userPwd"`

}

func aa() (m *Monster){
	print(m)
	fmt.Println()
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