package utils

import (
	"chapterDemo/chapter18/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (tf *Transfer) WritePkg(data []byte) (err error) {
	var pakLen uint32
	pakLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(tf.Buf[:4], pakLen)
	//发送长度
	n, err := tf.Conn.Write(tf.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	//fmt.Printf("，发送消息的长度%d，内容=%s\n", len(data),string(data))
	//发送消息本身
	n, err = tf.Conn.Write(data)
	if n != int(pakLen) || err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}
	return
}

//将读取包的任务封装到了一个函数中 readPkg
func (tf *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取客户端发送的数据：")
	_, err = tf.Conn.Read(tf.Buf[:4])
	if err != nil {
		//err=errors.New("nihao")
		//fmt.Println("conn.Read err=",err)
		return
	}
	//fmt.Println("读取到的buf=",buf[:4])
	//根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(tf.Buf[0:4])
	//根据pkgLen读取消息
	n, err := tf.Conn.Read(tf.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	//将pkgLen反序列化成->message.Message
	err = json.Unmarshal(tf.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}
