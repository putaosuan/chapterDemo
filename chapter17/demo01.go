package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(n interface{}) {
	//通过反射获取传入的变量的type，king，值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(n)
	fmt.Println("tTyp=", rTyp)
	//2.获取到reflect.Value
	rVal := reflect.ValueOf(n)
	n2 := 2 + rVal.Int()
	fmt.Println(n2)
	fmt.Printf("rVal=%v,rVal类型=%T\n", rVal, rVal)
	//3.将rVal转换成interfa{}
	iV := rVal.Interface()
	//将interfa{}通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

//专门演示对结构体的反射
func reflectTest02(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Println("tTyp=", rTyp)
	//
	rVal := reflect.ValueOf(b)
	//获取变量对应的kind
	kind1 := rVal.Kind()
	kind2 := rTyp.Kind()
	fmt.Printf("%T,%T\n", kind1, kind2)

	iV := rVal.Interface()
	fmt.Printf("iV=%v,iV的类型=%T\n", iV, iV)
	//将interfa{}通过断言转换成需要的类型
	//使用前面学过的断言形式来更加灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu Name=%v\n", stu.Name)
	}

}

type Student struct {
	Name string
	Age  int
}
type Monster struct {
	Name string
	Age  int
}

func main() {
	//编写一个案例，对（基本数据类型、interface{}，reflect.Value）进行反射的基本操作
	//var num int = 100
	//reflectTest01(num)
	stu := Student{
		Name: "xiaomi",
		Age:  10,
	}
	reflectTest02(stu)
}
