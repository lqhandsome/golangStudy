package monster

import (
	"fmt"
	"testing"
)

func TestMonster_Store(t *testing.T) {
	monster := Monster{
		"红孩儿",
		600,
		"三味真火",
	}
	err := monster.Store()
	if err != nil {
		t.Fatalf("测试不通过%v=",err)
	}
	fmt.Println(monster)
	t.Log("测试通过")
}

func TestMonster_ReStore(t *testing.T) {
	//创建一个monster不指针值
	var monster Monster
	monster.Name = "红孩儿"
	err := monster.ReStore()
	if err != nil {
		t.Fatalf("测试失败=%v",err)
	}
	fmt.Println(monster)
	t.Log("通过")
}