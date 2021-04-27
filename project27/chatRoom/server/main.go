package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/project27/chatRoom/common/message"
	"io"
	"net"
)

func main() {
	listen,err := net.Listen("tcp","127.0.0.1:8889")
	if err != nil {
		fmt.Println("创建监听失败,",err)
		return
	}
	defer listen.Close()
	//提示信息
	fmt.Printf("服务器在 %v 监听...\n",listen.Addr())
	for {
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("accept失败",err)
		}
		//将连接信息发给协程
		go process(conn)
	}

}

func process(conn net.Conn){
	defer conn.Close()
	//读取客户端发送的信息

	for {
		msg,err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭了连接")
				return
			}
			fmt.Println("读取客户端数据失败readPkg.err=",err)
			return
		}
		fmt.Println("msg",msg)
		err = serverProcessMes(conn,&msg)
		if err != nil {
			fmt.Println("50errpr",err)
			return
		}
	}
}
func readPkg(conn net.Conn)(message message.Message,err error) {
	//获取数据

	buf := make([]byte,1024 * 8)
	fmt.Println("正在读取文件...")
	//conn.Read 是一个阻塞进程 必须发送端和接收端都不能关闭
	_, err = conn.Read(buf[:4])
	if  err != nil {
		fmt.Println("读取错误:",err)
		return
	}
	fmt.Println("读取到的buf=",buf[:4])

	//根据buf[0:4]转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	//根据pkgLen 读取消息内容
	n ,err := conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen && err != nil {
		fmt.Println("conn.Read fail err=",err)
		return
	}
	err = json.Unmarshal(buf[:pkgLen],&message)
	if err != nil {
		fmt.Println("Unmarshal fail err=",err)
		return
	}
	fmt.Println(string(buf[:pkgLen]))
	return
}

//处理用户登录
func serverProcessLogin(conn net.Conn,mes *message.Message)(err error) {
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
	writePkg(conn,data)
	return
}

//服务器发送给客户端信息
func writePkg(conn net.Conn,data []byte)(err error) {
	//先发送一个长度给对方
	var pakLen uint32
	pakLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pakLen)
	//发送长度
	n,err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}

	//发送data
	n, err = conn.Write(data)
	if err != nil || n != int(pakLen){
		fmt.Println("写入data失败",err)
		return
	}
	fmt.Println("发送给客户端的数据：",string(data))
	return
}

//根据客户端发送消息类型不同调用不同的函数
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		serverProcessLogin(conn,mes)
			//处理登录
	case message.RegisterMesType:
			//处理注册
	default:
		errors.New("不符合的消息类型")
	}
	return nil
}