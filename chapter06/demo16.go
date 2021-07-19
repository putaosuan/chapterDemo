package main

import (
	"fmt"
	"strings"
)

func main() {
	//8.查找字串是否在指定的字符串中
	b := strings.Contains("seafood", "sea")
	fmt.Println(b)
	//9.统计一个字符串有几个指定的字符串
	num := strings.Count("hello", "l")
	fmt.Println(num)
	//10.不区分大小写字符串比较
	b = strings.EqualFold("abc", "Abc")
	fmt.Println(b)
	//11.返回字串在字符串中最后一次出现的index，如没有，返回-1
	index := strings.Index("go golang", "go")
	fmt.Println(index)
	//12.返回字串在字符串中最后一次出现的index，如没有，返回-1
	index = strings.LastIndex("go golang", "go")
	fmt.Println(index)
	//13.将指定的字符串替换成另一个字符串：n表示替换几个，-1表示全部替换
	str2 := "go golang"
	str := strings.Replace(str2, "g", "d", -1)
	fmt.Println(str)
	//14.按照某个指定的字符，为分割字符，将一个字符串拆成字符数组
	strArr := strings.Split("hello,go,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Println(strArr[i])
	}
	//15.将字符串进行大小写转换
	str = "hello Go"
	str = strings.ToLower(str)
	fmt.Println(str)
	str = strings.ToUpper(str)
	fmt.Println(str)
	//16/将字符串左右两边的空格去掉
	str = strings.TrimSpace(" tn a lone gopher ntm ")
	fmt.Println(str)
	//17.将字符串左右两边的指定字符去掉
	str = strings.Trim("!hello!", "!")
	fmt.Println(str)
	//18.将左边指定的字符去掉
	//19.将右边指定的字符去掉
	//20.判断字符串是否以指定的字符串开头
	b = strings.HasPrefix("ftp://192.168.1.0", "ftp")
	fmt.Println(b)
	//20.判断字符串是否以指定的字符串结尾
	b = strings.HasSuffix("gogogodd", "dd")
	fmt.Println(b)
}
