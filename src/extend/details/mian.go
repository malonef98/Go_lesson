package main

import "fmt"

type A struct {
	Name string
	age int
}

func (a *A) SayOk()  {
	fmt.Println("A SayOk",a.Name)
}

func (a *A) hello()  {
	fmt.Println("A hello",a.Name)
}

type B struct {
	Name string
	age int
	A
}

func (b *B) hello ()  {
	fmt.Println("B hello",b.Name)
}

func main() {
	var b B
	//字段的赋值
	b.Name = "tom"
	b.A.Name = "jack"
	b.age = 19
	//同名方法的调用
	b.hello()
	b.A.hello()
}

