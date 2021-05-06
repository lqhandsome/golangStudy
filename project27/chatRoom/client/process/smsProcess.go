package process

import (
	"encoding/json"
	"fmt"
	"go_code/project27/chatRoom/client/utils"
	"go_code/project27/chatRoom/common/message"
)

type SmsProcess struct {
}

//发送群聊的消息
func (this *SmsProcess) SendGroupMes(context string) (err error) {

	//创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//创建SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = context
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json失败", err.Error())
		return err
	}
	mes.Data = string(data)
	//整体序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json失败", err.Error())
		return err
	}

	//发送数据
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	if err = tf.WritePkg(data);err != nil {
		fmt.Println("发送失败")
		return
	}
	fmt.Println("发送成功")
	return
}
