package main

import "fmt"

func main() {
	exp := "3+2*6-2"
	ch := exp[0:1]
	temp := int(ch[0])
	fmt.Printf("%T,%v\n", ch, ch)
	fmt.Printf("%T,%v\n", temp, temp)
	ch1 := exp[0]
	fmt.Printf("%T,%v\n", ch1, ch1)

	i := 1
	if i > 0 {
		if i != 2 {
			i = 2
		} else {
			fmt.Println("你好")
		}
	}

}
