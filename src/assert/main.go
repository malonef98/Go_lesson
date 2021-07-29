package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	//空接口可以接受任何变量
	var a interface{}
	var point Point = Point{1,2}
	a = point
	// 如何将a赋值给一个Point变量
	var b Point
	//b = a   不可以
	//b = a.(Point)  可以   类型断言
	b = a.(Point)
	fmt.Println(b)


	//类型断言其他类型(带检测)
	var x interface{}
	// 空接口，可以接受任何类型
	var b2 float32 = 1.1
	x = b2
	// x => float32 [使用类型断言]
	if y , ok := x.(float32); ok == true {
		fmt.Println("转换成功")
		fmt.Printf("y 的类型是%T，值是%v",y,y)
	}else {
		fmt.Println("convert fail")
	}

	fmt.Println("继续执行")

}
