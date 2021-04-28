package process

import (
	"encoding/json"
	"fmt"
	"go_code/project27/chatRoom/common/message"
	"go_code/project27/chatRoom/server/utils"
	"net"
)

type UserProcess struct {
	//字段
	Conn net.Conn
}

//处理用户登录
func (this *UserProcess)ServerProcessLogin(mes *message.Message)(err error) {
	//1.先从mes.data中取出数据并反序列化
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)

	if err != nil {
		fmt.Println("格式化mes.data失败")
		return
	}

	//申明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//声明一个loginResMes
	var loginResMes message.LoginResMes
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "账号或者密码不正确，请重新登录"
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化失败108",err)
	}

	//将data赋值给resMes
	resMes.Data = string(data)

	//对要返回的数据进行整体序列化

	data,err =json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败119",err)
	}
	
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	tf.WritePkg(data)
	return
}