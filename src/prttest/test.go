package main

import "fmt"

func main() {
	//1 声明变量
	var age int = 18
	var age2 int =  20
	//声明指针类型 *int  指向int类型的指针
	var prt *int = &age
	var prt2 *int = &age2
	//改变变量地址的值
	*prt = 20
	*prt2 = 18
	fmt.Println(age , age2)
}