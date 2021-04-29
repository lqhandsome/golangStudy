package process

import (
	"encoding/json"
	"fmt"
	"go_code/project27/chatRoom/common/message"
	"go_code/project27/chatRoom/server/model"
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

	user, err := model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)

	if err != nil {
		loginResMes.Code = 500

		if err == model.ERROR_USER_NOTEXISTS  {
			loginResMes.Error = "用户不存在"
		}
		if err == model.ERROR_USER_PWD {
			loginResMes.Error = "密码不正确"
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user.UserId,"登录成功")
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