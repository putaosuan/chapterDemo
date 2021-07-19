package main

import (
	"errors"
	"fmt"
	"strconv"
)

//使用数组来模拟一个栈
type Stack struct {
	MaxTop int
	Top    int
	arr    [5]int //数组模拟栈
}

func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop-1 {
		fmt.Println("栈满")
		return errors.New("stack full")
	}
	this.Top++
	//放入数据
	this.arr[this.Top] = val
	return
}

//遍历栈
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("栈空")
		return
	}
	fmt.Println("栈的情况如下")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("栈空")
		return 0, errors.New("stack empty")
	}
	//先取值
	val = this.arr[this.Top]
	this.Top--
	return val, nil

}
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		return 1
	} else if oper == 43 || oper == 45 {
		return 0
	}
	return res
}
func (this *Stack) Cal(num1, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}
func main() {
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "30+2*6-20"
	//定义一个index，帮助扫描exp
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keeNum := ""
	for {
		ch := exp[index : index+1]
		temp := int(ch[0])
		if operStack.IsOper(temp) {
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					//fmt.Println("1:",result)
					numStack.Push(result)
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}

		} else {
			keeNum += ch
			if index+1 == len(exp) {
				val, _ := strconv.ParseInt(keeNum, 10, 64)
				numStack.Push(int(val))
			} else {
				if operStack.IsOper(int(exp[index+1 : index+2][0])) {
					val, _ := strconv.ParseInt(keeNum, 10, 64)
					numStack.Push(int(val))
					keeNum = ""
				}
			}

			//val,_:=strconv.ParseInt(ch,10,64)
			//numStack.Push(int(val))
		}
		if index+1 == len(exp) {
			break
		}
		index++
	}
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		//fmt.Println(result)
		numStack.Push(result)
	}
	//num1,_=numStack.Pop()
	//num2,_=numStack.Pop()
	//oper,_=operStack.Pop()
	//result=operStack.Cal(num1,num2,oper)
	//numStack.Push(result)

	res, _ := numStack.Pop()
	fmt.Println(res)

}
