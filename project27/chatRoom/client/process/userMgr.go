package process

import (
	"fmt"
	"go_code/project27/chatRoom/client/model"
	"go_code/project27/chatRoom/common/message"
)

//客户端要维护的map
var onlineUsers map[int]*message.User  = make(map[int]*message.User,10)
var CurUser model.CurUser //在登陆完成后对其初始化
//在客户端显示当前在线的用户
func outputOnlineUser() {
	fmt.Println("当前用户列表")
	//遍历所有用户
	for id := range onlineUsers {
		fmt.Println("用户id-user",id)
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	//如果没有
	if !ok {
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus =  notifyUserStatusMes.UserStatus
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}