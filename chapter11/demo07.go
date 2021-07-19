package main

import "fmt"

type Usb7 interface {
	Say()
}
type Stu struct {
}

func (s *Stu) Say() {
	fmt.Println("Say()")
}

func main() {
	var stu Stu = Stu{}
	var u Usb7 = &stu
	u.Say()
	fmt.Println(u)
}
