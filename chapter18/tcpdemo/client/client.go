package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client err=", err)
		return
	}
	defer conn.Close()
	//fmt.Println("conn 成功",conn)
	//功能一：客户端可以发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin) //os.Stdin代表标准输入（终端）
	//从终端读取一行用户输入，并准备发送给服务器
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		//需要将\n去掉
		line = strings.Trim(line, "\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}
		//再将line 发送给服务器
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn write err=", err)
		}
		fmt.Printf("客户端发送了%d字节的数据,并退出\n", n)
	}

}
