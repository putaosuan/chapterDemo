package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//统计一个文件里面含有的英文、数字、空格以及其他字符数量
type CharCount struct {
	ChCount    int //记录英文个数
	NumCount   int //记录数字的个数
	SpaceCount int //记录空格的个数
	OtherCOunt int //记录其他字符的个数
}

//定义一个结构体，用于保存统计结果
func main() {
	//打开一个文件，创建一个reader
	//没读取一行，就统计该行有多少英文，数字和其他字符
	fileName := "d:/test1/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("file open err = ", err)
	}
	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')

		for _, v := range str {
			fmt.Printf("%c", v)
			switch {
			case 'a' <= v && v <= 'z', 'A' <= v && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v > '0' && v < '9':
				count.NumCount++
			default:
				count.OtherCOunt++
			}
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("字符的个数：%v,空格的个数：%v，数字：%v,其他字符 ：%v\n", count.ChCount,
		count.SpaceCount, count.NumCount, count.OtherCOunt)
}
