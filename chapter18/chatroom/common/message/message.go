package message

const (
	LoginMesType            = "Message"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterMesType"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的内容
}

//定义两个消息，后面需要再增加
type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
type LoginResMes struct {
	Code   int    `json:"code"` //返回状态码，500表示该用户未注册，200表示登录成功
	UserId []int  //增加字段，保存用户切片
	Error  string `json:"error"`
}
type RegisterMes struct {
	User User `json:"user"`
}
type RegisterResMes struct {
	Code  int //返回状态码 400表示已经占有，200表示登陆成功
	Error string
}

//为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

//增加一个SmsMes //发送消息
type SmsMes struct {
	Content string `json:"content"` //内容
	User
}
