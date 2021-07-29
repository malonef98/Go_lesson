package demo2

import (
	"testing"
) //go的testing测试框架

//编写一个测试用例  去测试函数
func TestAddUpper(t *testing.T)  {
	//调用
	res := AddUpper(10)
	if res != 10{
		t.Fatalf("错咯")
	}

	t.Logf("对咯")
}
