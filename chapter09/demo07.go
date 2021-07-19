package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称"
	}
}
func main() {
	user := make(map[string]map[string]string)
	user["smith"] = make(map[string]string)
	user["smith"]["pwd"] = "666666"
	user["smith"]["nichname"] = "nicheng"
	modifyUser(user, "mary")
	modifyUser(user, "switch")
	modifyUser(user, "smith")
	fmt.Println(user)
}
