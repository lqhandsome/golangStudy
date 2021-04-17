package cal

import (
	"testing"
)


func TestAddUpper(t *testing.T) {
	//调用
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("upper（10）函数执行错误期望值=%v,实际值=%v\n",55,res)
	}

	t.Logf("upper(10)执行正确")
}
func TestUpper(t *testing.T) {
	//调用
	res := addUpper(10)
	if res != 51 {
		t.Fatalf("upper（10）函数执行错误期望值=%v,实际值=%v\n",55,res)
	}

	t.Logf("upper(10)执行正确")
}
