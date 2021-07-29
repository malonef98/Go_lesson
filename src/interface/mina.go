package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {

}

func (p Phone) Start()  {
	fmt.Println("手机开始工作......")
}

func (p Phone) Stop()  {
	fmt.Println("手机结束工作......")
}

type Camera struct {

}

func (camera Camera) Start()  {
	fmt.Println("相机开始工作......")
}

func (camera Camera) Stop()  {
	fmt.Println("相机结束工作......")
}

type Computer struct {

}
//编写一个Working方法 接收一个Usb接口类型变量
//只要实现Usb接口，就要实现接口中声明的所有方法
func (computer Computer) Working(usb Usb)  {
	//通过接口调用接口内方法
	usb.Start()
	usb.Stop()
}


func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)
}