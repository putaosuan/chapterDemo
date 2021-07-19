package main

import (
	"fmt"
	"strconv"
)

func main() {
	//2.字符串遍历
	str2 := "hello GO语言"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c \t", r[i])
	}
	fmt.Println()
	//3.字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误")
	} else {
		fmt.Println("转换的结果是：", n)
	}
	//4.整数转字符串
	str := strconv.Itoa(123456)
	fmt.Println(str)
	//5.字符串转[]byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)
	//6.[]byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Println(str)
	//7.10进制转2，6，8进制
	str = strconv.FormatInt(123, 2)
	fmt.Println(str)
	str = strconv.FormatInt(123, 8)
	fmt.Println(str)
	str = strconv.FormatInt(123, 16)
	fmt.Println(str)

}
