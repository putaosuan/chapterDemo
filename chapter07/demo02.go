package main

import "fmt"

func main() {
	var myChars [26]byte
	for i := 0; i < len(myChars); i++ {
		myChars[i] = 'A' + byte(i)
	}
	for _, v := range myChars {
		fmt.Printf("%c\t", v)
	}
}
