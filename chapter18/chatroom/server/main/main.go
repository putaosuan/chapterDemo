package main

import (
	"chapterDemo/chapter18/chatroom/server/model"
	"fmt"
	"net"
	"time"
)

////编写一个ServerProcessMes函数
////功能：根据客户端发送消息种类的不同，决定调用哪个函数来处理
//func ServerProcessMes(conn net.Conn,mes *message.Message)(err error)  {
//	switch mes.Type {
//		case message.LoginMesType:
//			//处理登录函数
//			serverProcessLogin(conn,mes)
//		case message.RegisterMesType:
//			//处理注册函数
//		default:
//			fmt.Println("消息类型不存在，无法处理")
//	}
//	return
//}
////编写一个函数serverProcessLogin,专门处理登录请求
//func serverProcessLogin(conn net.Conn,mes *message.Message)(err error)  {
//	//1.先从mes.Data,并直接反序列化成LoginMes
//	var loginMes message.LoginMes
//	err=json.Unmarshal([]byte(mes.Data),&loginMes)
//	if err!=nil{
//		fmt.Println("json.Unmarshal err=",err)
//		return
//	}
//	//1.先声明一个resMes
//	var resMes message.Message
//	resMes.Type=message.LoginResMesType
//	//2.声明一个loginResMes，并赋值
//	var loginResMes message.LoginResMes
//
//	if loginMes.UserId==100&&loginMes.UserPwd=="123456"{
//		loginResMes.Code=200
//	}else {
//		loginResMes.Code=500
//		loginResMes.Error="该用户不存在，请注册再使用"
//	}
//	//3.将loginResMes序列化
//	data,err:=json.Marshal(loginResMes)
//	if err!=nil{
//		fmt.Println("json.Marshal err=",err)
//		return
//	}
//	//4.将data赋值给resMes
//	resMes.Data=string(data)
//	//5.对resMes进行序列化，准备发送
//	data,err=json.Marshal(resMes)
//	if err!=nil{
//		fmt.Println("json.Marshal err=",err)
//		return
//	}
//	//6.发送data，我们将其封装到writePkg函数中
//	err=writePkg(conn,data)
//	return
//}
//func writePkg(conn net.Conn,data []byte) (err error) {
//	var pakLen uint32
//	pakLen= uint32(len(data))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[:4],pakLen)
//	//发送长度
//	n,err:=conn.Write(buf[:4])
//	if n!=4||err!=nil{
//		fmt.Println("conn.Write err=",err)
//		return
//	}
//	//fmt.Printf("，发送消息的长度%d，内容=%s\n", len(data),string(data))
//	//发送消息本身
//	n,err=conn.Write(data)
//	if n!=int(pakLen)||err!=nil{
//		fmt.Println("conn.Write(data) err=",err)
//		return
//	}
//	return
//}
////将读取包的任务封装到了一个函数中 readPkg
//func readPkg(conn net.Conn)(mes message.Message,err error)  {
//	buf:=make([]byte,8096)
//	fmt.Println("读取客户端发送的数据：")
//	_,err=conn.Read(buf[:4])
//	if err!=nil{
//		//err=errors.New("nihao")
//		//fmt.Println("conn.Read err=",err)
//		return
//	}
//	//fmt.Println("读取到的buf=",buf[:4])
//	//根据buf[:4]转成一个uint32类型
//	var pkgLen uint32
//	pkgLen=binary.BigEndian.Uint32(buf[0:4])
//	//根据pkgLen读取消息
//	n,err:=conn.Read(buf[:pkgLen])
//	if n!=int(pkgLen)||err!=nil{
//		return
//	}
//	//将pkgLen反序列化成->message.Message
//	err=json.Unmarshal(buf[:pkgLen],&mes)
//	if err!=nil{
//		fmt.Println("json.Unmarshal err=",err)
//		return
//	}
//	return
//}

func process(conn net.Conn) {
	defer conn.Close()

	//for{
	//	mes,err:=readPkg(conn)
	//	if err!=nil{
	//		if err==io.EOF{
	//			fmt.Println("客户端退出，服务器端也退出")
	//			return
	//		}else {
	//			fmt.Println("readPkg err=",err)
	//			return
	//		}
	//	}
	//	//fmt.Println("mes=",mes)
	//	ServerProcessMes(conn,&mes)
	//
	//}
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误,err=", err)
		return
	}
}
func initUserDao() {
	//这里pool本身就是一个全局变量
	//这里需要注意一个初始化顺序
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	//当启动服务时，我们就去初始化我们的连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() err=", err)
		}
		//一旦连接成功，则启动一个协程和客户端保持通信
		go process(conn)
	}
}
