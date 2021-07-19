package testcase2

import (
	"fmt"
	"testing"
)

func TestStore(t *testing.T) {
	var monster = &Monster{
		Name:  "小米",
		Age:   11,
		Skill: "手机制造",
	}
	var b bool
	b = monster.Store()
	if b == false {
		t.Fatalf("失败了,期望值%v,实际值%v", true, b)
	}
	t.Logf("成功了")
}
func TestReStore(t *testing.T) {
	var monster = &Monster{}
	fmt.Println(monster)
	var b bool
	b = monster.ReStore()
	if b == false {
		t.Fatalf("失败了,期望值%v,实际值%v", true, b)
	}
	if monster.Name == "小米" {
		fmt.Println("你好")
	}
	t.Logf("成功了")
}
