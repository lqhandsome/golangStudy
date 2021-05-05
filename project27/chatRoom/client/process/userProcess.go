package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/project27/chatRoom/common/message"
	"go_code/project27/chatRoom/server/utils"
	"net"
	"os"
)

type UserProcess struct {

}

func (this *UserProcess)Login(userId int,pwd string)(err error) {
	//连接服务端
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接失败：",err)
		return
	}
	//延时关闭连接
	defer conn.Close()

	//准备发送信息
	var mes message.Message
	mes.Type = message.LoginMesType

	//创建一个LoginMes结构体
	var LoginMes message.LoginMes
	LoginMes.UserId = userId
	LoginMes.UserPwd = pwd

	//将数据LoginMes json格式化
	jsonData,err := json.Marshal(LoginMes)
	if err != nil{
		fmt.Println("LoginMes格式化失败",err)
		return
	}
	//将格式化的数据转为字符复制给data
	mes.Data = string(jsonData)

	//将整个mes json格式化
	data,err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mes格式化失败",err)
		return
	}

	dataLen := len(data)
	var pakLen uint32
	pakLen = uint32(dataLen)
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pakLen)
	n,err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	fmt.Println("客户端发送的长度",data,buf[0:4])
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("写入data失败",err)
		return
	}
	tf := utils.Transfer{
		Conn: conn,
	}
	msg,err := tf.ReadPkg()
	if err != nil {
		fmt.Println("读取失败76",err)
		return
	}

	//将mes的data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(msg.Data),&loginResMes)
	fmt.Println("用户端读取的loginResMes",loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
		//1.显示登陆成功后的菜单
		fmt.Println("当前在线用户的列表如下：")
		for _,  v := range loginResMes.UserIds {
			if v == userId {
				continue
			}
			fmt.Println("用户id：\t",v)
			//完成客户端维护的map初始化
			user := &message.User{
				UserId: v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		//开启协程，保持客户端与服务端保持通讯
		go serverProcessMes(conn)
		ShowMenu()

	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (this *UserProcess) RegisterUser(userId int,userPwd string,userName string)(err error){
	//1.连接服务端
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接失败：",err)
		return
	}
	//延时关闭连接
	defer conn.Close()
	//2.创建一个message.Message结构体用来发送数据
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3.创建一个RegisterMes结构体原来保存用户输入的数据
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	//4.将用户输入的数据格式化
	data,err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("格式化用户数据错误,",err)
		return
	}
	//5.将用户输入的数据格式化存入data中
	mes.Data = string(data)

	//6.讲要发送的数据整体格式化
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("格式化mes错误,",err)
		return
	}

	//7.发送数据
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息出错，",err)
		return
	}
	//8读取客户端发送的数据
	msg,err := tf.ReadPkg()
	fmt.Println("137",msg)
	if err != nil {
		fmt.Println("接收信息出错，",err)
		return
	}
	//将mes的data部分反序列化成LoginResMes
	var RegisterResMes message.RegisterResMes
	err = json.Unmarshal([]byte(msg.Data),&RegisterResMes)
	fmt.Println("用户端读取的loginResMes",RegisterResMes)
	if RegisterResMes.Code == 200 {
		fmt.Println("注册成功")
		os.Exit(0)
		//1.显示登陆成功后的菜单
		//go serverProcessMes(conn)
		//ShowMenu()

	} else  {
		fmt.Println(RegisterResMes.Error)
	}
	os.Exit(0)
	return
}