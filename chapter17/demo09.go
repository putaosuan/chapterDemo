package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Printf("%v完成了运算，%v-%v=%v\n", name, c.Num1, c.Num2, c.Num1-c.Num2)
}
func reflect9(b interface{}) {
	rtyp := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)
	rVal = rVal.Elem()
	rtyp = rtyp.Elem()
	rVal.FieldByName("Num1").SetInt(8)
	rVal.FieldByName("Num2").SetInt(3)
	num := rVal.NumField()
	for i := 0; i < num; i++ {
		fmt.Printf("Field%v的值：%v\n", i, rtyp.Field(i).Name)
		fmt.Printf("Field%v的值：%v\n", i, rVal.Field(i))
	}

	var p []reflect.Value
	p = append(p, reflect.ValueOf("tom"))
	rVal.Method(0).Call(p)
}
func main() {
	c := Cal{}
	reflect9(&c)
}
