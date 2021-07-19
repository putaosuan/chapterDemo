package main

import "fmt"

func main() {
	//我们要存放三个学生的信息，每个学生有name和sex
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "北京长安街"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu01"])
	fmt.Println(studentMap["stu01"]["address"])

	//遍历
	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			//fmt.Printf("studentMap[%v][%v]=%v\n",k1,k2,v2)
			fmt.Printf("k2=%v,v2=%v\n", k2, v2)
		}
	}

}
