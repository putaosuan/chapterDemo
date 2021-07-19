package main

import (
	"fmt"
	"io"
	"net" //做网络开发时，net包含有我们需要的所有的方法和函数
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		//创建一个新切片
		buf := make([]byte, 1024)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有发送，那么协程就阻塞在这
		fmt.Printf("服务器在等待客户端%v发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("客户端退出err=", err)
			return
		}
		//3.显示客户端发送的信息到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}
func main() {
	fmt.Println("服务器开始监听...")
	//1.tcp 表示使用网络协议是tcp
	//2.0.0.0.0：8888表示在本地监听8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc conn=%v,客户端IP=%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
