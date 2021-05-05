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
	//ID
	UserId int

}

//通过所有在线的用户，有新用户上线
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	//遍历所有在线用户
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		_ = up.NotifyMeOnline(userId)
		//通知其他在线用用户
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) (err error){
		var mes message.Message
		mes.Type = message.NotifyUserStatusMesType

		var notifyUserStatusMes message.NotifyUserStatusMes
		notifyUserStatusMes.UserId = userId
		notifyUserStatusMes.UserStatus = message.UserOnline

		//将notifyUserStatusMes序列化

		data,err := json.Marshal(notifyUserStatusMes)
		if err != nil {
			fmt.Println("json格式化错误num44")
			return err
		}
		mes.Data = string(data)
		//将数据整体格式化
		data,err = json.Marshal(mes)
		if err != nil {
			fmt.Println("json格式化错误num44")
			return err
		}
		//创建TransFer实例发送数据
		tf :=  &utils.Transfer{
			Conn: this.Conn,

		}
	fmt.Println("aaaaaaaaaaaaaaaa",string(data),mes)
		//发送数据
		tf.WritePkg(data)

		return
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
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//登陆成功通知其他用户
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//讲在线用户放入切片返回给客户端
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds,id)
		}
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

func (this  *UserProcess) ServerProcessRegister(mes *message.Message)(err error) {
	//1.先从mes.data中取出数据并反序列化
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data),&registerMes)

	if err != nil {
		fmt.Println("格式化mes.data失败")
		return
	}

	//申明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterMesType

	//声明一个registerRes
	var registerResMes message.RegisterResMes

	//进入数据库添加用户
	err = model.MyUserDao.Register(registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	data,err := json.Marshal(registerMes)
	resMes.Data = string(data)

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	tf.WritePkg(data)
	//发送数据
	return
}