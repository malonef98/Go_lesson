package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start()  {
	fmt.Println("手机开始工作......")
}

func (p Phone) Stop()  {
	fmt.Println("手机结束工作......")
}

func (p Phone) Call()  {
	fmt.Println("打电话......")
}

type Camera struct {
	Name string
}

func (camera Camera) Start()  {
	fmt.Println("相机开始工作......")
}

func (camera Camera) Stop()  {
	fmt.Println("相机结束工作......")
}

type Computer struct {

}

func (computer Computer) Working(usb Usb)  {
	//通过接口调用接口内方法
	usb.Start()
	//如果Usb是指向Phone结构体变量，还需要执行Call方法
	//类型断言  [高级用法]
	if phone,ok := usb.(Phone); ok == true {
		phone.Call()
	}

	usb.Stop()
}


func main() {

	var usbArr = [3]Usb{}
	usbArr[0] = Phone{"VIVO"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"SONY"}

	var computer Computer
	for _,v := range usbArr{
		computer.Working(v)
	}
}