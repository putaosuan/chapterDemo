package message

//定义一个用户结构体
type User struct {
	//确定字段信息
	//为了序列化和反序列化成功，必须保证用户的json字符串的key和结构体对应字段tag名字一样
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
	Sex        string `json:"sex"`
}
