package message

const (
	LoginMesType 				= "LoginMes"
	LoginResMesType 			= "LoginResMes"
	RegisterMesType 			= "RegisterMesType"
	NotifyUserStatusMesType 	= "NotifyUserStatusMes"
	SmsMesType					=  "SmsMesType"

)

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)
type Message struct {
	Type string `json:"type"`//消息类型
	Data string `json:"data"`//数据
}

type LoginMes struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code int `json:"code"`//状态码
	Error string `json:"error"`//错误信息
	UserIds []int //在线用户的id
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code int `json:"code"`//状态码 400该用户已经占用，200表示注册成功
	Error string `json:"error"`//错误信息
}

//通知客户端有新用户上线
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	UserStatus int  `json:"userStatus"`
}

//增加一个SmsMes 发送的消息
type  SmsMes struct {
	Content string
	User//匿名结构体
}
