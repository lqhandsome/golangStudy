package process

import (
	"encoding/json"
	"fmt"
	"go_code/project27/chatRoom/common/message"
	"go_code/project27/chatRoom/server/utils"
	"net"
)

type SmsProcess struct {

}

//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message)(err error) {
	//保存数据
	var smsMes message.SmsMes
	//取出message的内容并反序列化
	err = json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("SendGroupMes反序列化失败",err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes序列化失败",err)
		return
	}
	//遍历所有在线的用户
	for id, user := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
			this.SendMesToEachOnlineUser(data,user.Conn)
	}
	return
}

//向单个用户发送消息
func (this *SmsProcess) SendMesToEachOnlineUser(date []byte,conn net.Conn) {

	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(date)
	if err != nil {
		fmt.Println("转发消息失败",err)
		return
	}
	fmt.Println("err54",err)
}