package main

import (
	"fmt"
	"go_code/project27/chatRoom/client/process"
)
var (
	userId int
	userPwd string
)
func main() {

		//接收用户的选择
	var key int
	//判断用户是否还继续显示菜单
	var loop = true
	for loop {
		fmt.Println("------------欢迎登陆多人聊天系统")
		fmt.Println("\t\t\t 1.登陆聊天室")
		fmt.Println("\t\t\t 2.注册用户")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("\t\t\t 请选择（1-3）")
		fmt.Scanf("%d\n", &key)
		switch key {
			case 1:
				fmt.Println("登陆聊天室")
				//用户要登录
				fmt.Print("请输入用户的Id:")
				fmt.Scanf("%d\n",&userId)
				fmt.Print("请输入用户密码:")
				fmt.Scanf("%v\n",&userPwd)
				up := &process.UserProcess{
					}
				_ =up.Login(userId,userPwd)
				loop = false
			case 2:
				fmt.Println("注册用户")
				loop = false
			case 3:
				fmt.Println("退出系统")
				loop = false
			default:
				fmt.Println("输入有误重新输入")
		}
	}
	if key == 1 {


	}
}
