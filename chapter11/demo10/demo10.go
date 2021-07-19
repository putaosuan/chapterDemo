package main

import "fmt"

type Usb interface {
	Start1()
	Stop1()
}
type Phone struct {
	Name string
}
type Camera struct {
	Name string
}

func (p Phone) Start1() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop1() {
	fmt.Println("手机结束工作")
}
func (p Phone) Call() {
	fmt.Println("打电话")
}
func (c Camera) Start1() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop1() {
	fmt.Println("相机结束工作")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start1()
	//如果指向Phone，使用call方法
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop1()
}
func main() {
	//定义一个接口，可以存放Phone和Camera的结构体变量
	var usb [3]Usb
	usb[0] = Phone{"你好"}
	usb[1] = Phone{"不知道"}
	usb[2] = Camera{"wobuhao"}

	fmt.Println(usb)
	var computer Computer
	for _, v := range usb {
		computer.Working(v)
	}
}
