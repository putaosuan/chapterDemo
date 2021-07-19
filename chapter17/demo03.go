package main

import (
	"fmt"
	"reflect"
)

func reflectTest03(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal.Type(), rVal.Kind())
	rV := rVal.Interface()
	f1, ok := rV.(float64)
	if ok {
		fmt.Println(f1)
	}

}
func main() {
	var v float64 = 1.2
	reflectTest03(v)
	fmt.Println("-----------")
	var str string = "hello"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("nihao")
	fmt.Printf("%v\n", str)
}
