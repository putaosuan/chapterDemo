package main

import "fmt"

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}
type TV struct {
	Goods
	Brand
}
type TV2 struct {
	*Goods
	*Brand
}

func main() {
	tv := TV{Goods{"电视机1", 5999}, Brand{"海尔", "青岛"}}
	tv2 := TV{
		Goods{
			Name:  "电视机002",
			Price: 4999,
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
	}
	fmt.Println(tv)
	fmt.Println(tv2)
	tv3 := TV2{&Goods{"电视机03", 6999}, &Brand{"创维", "河南"}}
	tv4 := TV2{
		&Goods{
			Name:  "电视机04",
			Price: 7999,
		},
		&Brand{
			Name:    "长虹",
			Address: "四川",
		},
	}
	fmt.Println(*tv3.Goods, *tv3.Brand)
	fmt.Println(tv4.Goods, tv4.Brand)
}
