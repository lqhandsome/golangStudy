package process

import (
	"encoding/json"
	"fmt"
	"go_code/project27/chatRoom/common/message"
)

func outputGroupMes(mes *message.Message) {
	//反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("outputGroupMes反序列化失败",err)
		return
	}

	//显示
	content := smsMes.Content
	info := fmt.Sprintf("用户id:\t%d 对大家说\t%s",smsMes.UserId,content)
	fmt.Println(info,"\n")
}
