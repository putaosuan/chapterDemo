package main

import "fmt"

func main() {
	var str string = "helloGO"
	b := []byte(str)
	b[5] = 'g'
	b[6] = 'o'
	fmt.Println(string(b))
	//转化成[]byte后可以处理英文和数字，但不能处理中文，
	var str2 string = "hello山东"
	r := []rune(str2)
	r[5] = '山'
	r[6] = '东'
	fmt.Println(string(r))

}
