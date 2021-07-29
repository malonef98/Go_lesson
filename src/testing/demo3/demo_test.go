package demo3

import "testing"

//测试用例  去测试函数Store
func TestStore(t *testing.T)  {
	//先创建一个Monster
	 monster := &Monster{
	 	Name: "PDD",
	 	Age: 30,
	 	Skill: "LOL",
	 }

	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误，希望为%v,实际为%v",true,res)
	}
	t.Fatalf("测试Store成功")
}

//测试用例  去测试函数Store
func TestReStore(t *testing.T)  {
	//先创建一个Monster
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.Store()错误，希望为%v,实际为%v",true,res)
	}
	t.Fatalf("测试ReStore成功")
}

