package main

import "fmt"

type Person3 struct {
	Name string
	Age  int
}

func main() {
	//方式二
	var p2 = Person3{"mary", 24}
	fmt.Println(p2)
	//方式三
	var p3 *Person3 = new(Person3)
	p3.Name = "smith"
	p3.Age = 30
	fmt.Println(*p3)
	//方式四
	var p4 *Person3 = &Person3{"xiaomi", 10}
	p4.Name = "小米"
	fmt.Println(*p4)

}
