package process

import (
	"encoding/json"
	"fmt"
	"go_code/project27/chatRoom/client/utils"
	"go_code/project27/chatRoom/common/message"
	"net"
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
	var content string

	//初始化SmsProcess实例
	smsProcess := &SmsProcess{
	}
	fmt.Scanf("%d\n",&key)
	switch key {
		case 1:
			outputOnlineUser()
			ShowMenu()
		case 2:
			fmt.Println("请输入要发送的信息")
			fmt.Scanf("%s\n",&content)
			err := smsProcess.SendGroupMes(content)
			if  err != nil{
				fmt.Println("发送失败")
				return
			}
			ShowMenu()
		case 4:
		fmt.Println("退出了系统")
			return
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
		switch msg.Type {
			//有人上线
			case message.NotifyUserStatusMesType :
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(msg.Data),&notifyUserStatusMes)
				updateUserStatus(&notifyUserStatusMes)

			//有人群发消息
			case message.SmsMesType :
				fmt.Println("客户端接收到的群发消息")
				fmt.Println(msg)
				outputGroupMes(&msg)
			default:
				fmt.Println("服务器返回了一个位置类型")
		}
		fmt.Println("msg=",msg)
	}
}