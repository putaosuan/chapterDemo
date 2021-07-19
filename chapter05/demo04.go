package main

import "fmt"

func main() {
	/*
		如果我们的字符串含有汉字，不能用传统的对字符串的遍历，因为传统时根据字节来进行遍历的,需要先转换成切片
	*/
	var str = "hello 北京！"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}
	for i, v := range str2 {
		fmt.Printf("%v,%c\n", i, v)
	}
}
