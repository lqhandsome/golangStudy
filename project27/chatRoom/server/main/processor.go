package main

import (
	"errors"
	"fmt"
	"go_code/project27/chatRoom/common/message"
	process2 "go_code/project27/chatRoom/server/process"
	"go_code/project27/chatRoom/server/utils"
	"io"
	"net"
)
//创建一个processor结构体
type Processor struct {
	Conn net.Conn
}
//根据客户端发送消息类型不同调用不同的函数
func (this *Processor)serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		up := &process2.UserProcess{
				this.Conn,
		}
		up.ServerProcessLogin(mes)
		//处理登录
	case message.RegisterMesType:
		//处理注册
	default:
		errors.New("不符合的消息类型")
	}
	return nil
}

func (this *Processor) mainProcess()(err error) {
	//读取客户端发送的信息

	for {
		//创建一个Transfer 实例完成读包的任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		msg,err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端关闭了连接")
				return err
			}
			fmt.Println("读取客户端数据失败readPkg.err=",err)
			return err
		}
		fmt.Println("msg",msg)
		err = this.serverProcessMes(&msg)
		if err != nil {
			fmt.Println("50errpr",err)
			return err
		}
		return err
	}
}