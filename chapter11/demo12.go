package main

import "fmt"

//编写一个函数，可以判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for i, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", i, v)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", i, v)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", i, v)
		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型，值是%v\n", i, v)
		case string:
			fmt.Printf("第%v个参数是字符串类型，值是%v\n", i, v)
		default:
			fmt.Println("不确定")

		}
	}
}
func main() {
	TypeJudge(1, 1.23, true, "北京")

}
