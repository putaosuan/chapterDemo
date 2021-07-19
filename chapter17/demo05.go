package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Monster5 struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float64
	Sex   string
}

func (m Monster5) GetNum(n1, n2 int) int {
	return n1 + n2
}
func (m Monster5) Set(name string, age int, score float64, sex string) {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Sex = sex
}
func (m Monster5) Print() {
	fmt.Println("---start---")
	fmt.Println(m)
	fmt.Println("---end---")
}
func TestStruct5(b interface{}) {
	//typ:=reflect.TypeOf(b)
	val := reflect.ValueOf(b)
	//获取对应的类型
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("大象")
	fmt.Printf("struct has %d fileds \n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Elem().Field(i).Kind())
	}
	//获取到该结构体有多少个方法,方法的排序默认是按照 函数名（ASCLL）
	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d method\n", numOfMethod)
	val.Elem().Method(1).Call(nil)
	//
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Elem().Method(0).Call(params) //传入的参数是[]reflect.Value,返回的是[]reflect.Value
	fmt.Printf("%v\n", res[0].Int())

}
func main() {
	m := Monster5{
		Name:  "小米",
		Age:   10,
		Score: 66.8,
	}

	//Marsal就是通过反射获取到struct的tag值
	result, _ := json.Marshal(m)
	fmt.Println("json result:", string(result))
	TestStruct5(&m)
	fmt.Println(m)
}
