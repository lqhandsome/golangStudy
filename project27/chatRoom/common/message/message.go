package message

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMesType"
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
}

type RegisterMes struct {

}


