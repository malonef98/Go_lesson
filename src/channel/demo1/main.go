package main

import "fmt"

func main() {
	// 1 创建一个可以存放三个int类型的管道
	var intChan chan int
	intChan = make(chan int ,3)

	// 2 intChan是什么
	fmt.Printf("intchan的值%v，实际地址%v",intChan,&intChan)

	// 3 向管道写入数据
	// <- 写入符
	intChan<- 10
	num := 211
	intChan<- num
	//当我们给管道写入数据的时候  不能超过其的容量

	// 4 查看管道长度和cap容量
	fmt.Printf("\nchannel len = %v,cap = %v",len(intChan),cap(intChan))

	// 5 从管道中读取数据
	var num2 int
	num2 = <- intChan
	fmt.Println("\nnum2",num2)
	//取出数据后管道长度会变  容量不会变
	fmt.Printf("\nchannel len = %v,cap = %v",len(intChan),cap(intChan))

	// 6 在没有使用协程的情况下  如果通道中数据取完 就会报错

}
