package model

//定义一个用户的结构体

type User struct {
	UserId int `json:"userId"`
	UserName string `json:"userName"`
	UserPwd string `json:"userPwd"`
}
