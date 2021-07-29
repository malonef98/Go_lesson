package main

import "fmt"

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

type Usb interface {
	Start()
	Stop()
}

func (p Phone) Start()  {
	fmt.Println("手机开始工作...")
}

func (p Phone) Stop()  {
	fmt.Println("手机停止工作...")
}


func (c Camera) Start()  {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop()  {
	fmt.Println("相机停止工作...")
}


func main() {
	//定义一个Usb接口数组，可以存放实现接口的Phone和Camera结构体
	//这就体现出多态数据
	var usbArr [3]Usb
	usbArr[0] = Phone{"Vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"Sony"}

	fmt.Println(usbArr)

}
