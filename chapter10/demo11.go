package main

import "fmt"

type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (c *Calcuator) GetRes(operator byte) float64 {
	switch operator {
	case '+':
		return c.Num1 + c.Num2
	case '-':
		return c.Num1 - c.Num2
	case '*':
		return c.Num1 * c.Num2
	case '/':
		return c.Num1 / c.Num2
	default:
		fmt.Println("输入错误")
		return 0
	}
}
func main() {
	var c Calcuator
	c.Num1 = 4.0
	c.Num2 = 3.0
	res := c.GetRes('*')
	fmt.Println(res)
}
