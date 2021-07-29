package main

import "fmt"

func main() {
	intChan := make(chan int ,3)
	intChan <- 10
	intChan <- 20


	//关闭通道
	close(intChan)

	//这时候不能够在写入
	//intChan <- 300

	//读取数据是可以的
	n1 := <- intChan
	fmt.Println(n1)


	//遍历管道
	intChan2 := make(chan int,100)
	//channel 中放入100个数据
	for i := 0 ; i < 100 ; i++ {
		intChan2 <- i*2
	}

	//关闭管道
	close(intChan2)
	for n := range intChan2{
		fmt.Println("管道中取出的值=",n)
	}

}
