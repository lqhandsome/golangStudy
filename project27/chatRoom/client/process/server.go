package process

import (
	"fmt"
	"go_code/project27/chatRoom/server/utils"
	"net"
	"os"
)

//显示登陆成功后的界面

func ShowMenu() {
	fmt.Println("--------恭喜登陆成功-------")
	fmt.Println("--------1. 显示在线用户列表-------")
	fmt.Println("--------2. 发送消息-------")
	fmt.Println("--------3. 信息列表-------")
	fmt.Println("--------4. 退出系统-------")
	fmt.Println("请选择（1-4）：")
	var key int
	fmt.Scanf("%d\n",&key)
	switch key {
		case 1:
		fmt.Println("111")
		case 2:
		fmt.Println("111")
		case 4:
		fmt.Println("退出了系统")
		os.Exit(200)
		default:
		fmt.Println("输入错误")
	}
}

//和服务器端保持通信
func serverProcessMes(conn net.Conn) {
	//创建一个transFer实例
	//1.
	tf := utils.Transfer{
		Conn: conn,
	}
	for  {
		fmt.Println("客户端正在等待服务器发送的消息")
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=",err)
			return
		}
		fmt.Println("msg=",msg)
	}
}