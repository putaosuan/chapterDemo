package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := runtime.NumCPU()
	fmt.Println(num)
	//可以自己设置使用多个CPU
	runtime.GOMAXPROCS(num)

}
