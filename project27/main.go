package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

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