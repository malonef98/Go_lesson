package main

import "fmt"

type Monkey struct {
	Name string
}

func (m *Monkey) climbing()  {
	fmt.Println("生来会爬树...")
}

// 声明接口
type BirdAble interface {
	Flying()
}

// 声明接口
type FishAble interface {
	Swim()
}

type LittleMonkey struct {
	//继承
	Monkey
}

//实现接口
func (l LittleMonkey) Flying() {
	fmt.Println("通过学习小猴子会飞...")
}

func (l LittleMonkey) Swim() {
	fmt.Println("通过学习小猴子会游泳...")
}

func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "littlemonkey",
		},
	}

	monkey.climbing()
	monkey.Flying()
	monkey.Swim()
}