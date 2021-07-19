package main

import "fmt"

//声明/定义一个接口
type Usb interface {
	Start()
	Stop()
}
type Phone struct {
}

//让Phone实现Usb的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

//编写一个Working方法，接受一个Usb接口类型的变量
//只要实现了Usb接口
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}
func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(phone)
	computer.Working(camera)
}
