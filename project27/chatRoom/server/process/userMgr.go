package process

import "errors"

var (
	userMgr  *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess,1024),
	}
}


//增加一個用戶鏈接
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//刪除一個用戶
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers,userId)
}

//获取所有在线用户
func (this *UserMgr) GetAllOnlineUsers ()  map[int]*UserProcess{
	return this.onlineUsers
}

//根据id返回指定的用户
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess , err error) {
	up ,ok :=  this.onlineUsers[userId]
	if !ok {
		err = errors.New("用户不在线")
		return up,err
	}
	return
}