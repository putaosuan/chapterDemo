package main

//func login(userId int,userPwd string) (err error) {
//	//fmt.Printf("userId = %d userPwd = %s\n",userId,userPwd)
//	//return nil
//	//1.连接到服务器
//	conn,err:=net.Dial("tcp","localhost:8889")
//	if err!=nil{
//		fmt.Println("ner.Dial err=",err)
//		return
//	}
//	defer conn.Close()
//	//2.准备通过conn发送消息给服务
//	var mes message.Message
//	mes.Type=message.LoginMesType
//	//3.创建一个LoginMes结构体
//	var loginMes message.LoginMes
//	loginMes.UserId=userId
//	loginMes.UserPwd=userPwd
//	//4.将loginMes序列化
//	data,err:=json.Marshal(loginMes)
//	if err!=nil{
//		fmt.Println("son.Marshal err=",err)
//		return
//	}
//	//5.把mes.Data赋给
//	mes.Data=string(data)
//	//6.将mes序列化
//	data,err=json.Marshal(mes)
//	if err!=nil{
//		fmt.Println("son.Marshal err=",err)
//		return
//	}
//	//7.这个时候，data就是我们要发送的数据
//	//先获取到一个长度->转换成一个表示长度的byte切片
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
//	//fmt.Printf("客户端，发送消息的长度%d，内容=%s\n", len(data),string(data))
//	//发送消息本身
//	_,err=conn.Write(data)
//	if err!=nil{
//		fmt.Println("conn.Write(data) err=",err)
//		return
//	}
//	//休眠
//	//time.Sleep(10*time.Second)
//	//fmt.Println("休眠了一会")
//	//这里还需要处理服务器端返回的信息
//	mes,err=readPkg(conn)
//	if err!=nil{
//		fmt.Println("readPkg err=",err)
//		return
//	}
//	//将mes的Data部分发序列化成LgoinResMes
//	var loginResMes message.LoginResMes
//	err=json.Unmarshal([]byte(mes.Data),&loginResMes)
//	if loginResMes.Code==200{
//		fmt.Println("登录成功")
//	}else if loginResMes.Code==500{
//		fmt.Println(loginResMes.Error)
//	}
//	return
//
//}
