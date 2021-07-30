package main

import (
	"fmt"
	"time"
)

func writeDate(intChan chan int){
	for i := 0 ; i < 50 ; i++ {
		//放入数据
		intChan <- i
		fmt.Println("writeData ",i)
	}
	time.Sleep(time.Second * 5)
	//关闭
	close(intChan)
}

//读取数据
func readData(intChan chan int,exitChan chan bool){
	for {
		v , ok := <- intChan
		if !ok {
			break
		}
		fmt.Println("读取到数据=",v)
	}
	//读取完数据后  任务完成
	exitChan <- true
	close(exitChan)
}


func main() {

	//创建两个管道
	intChan := make(chan int , 50)
	exitChan := make(chan bool, 1)

	//创建两个协程
	go writeDate(intChan)
	go readData(intChan,exitChan)

	//程序延时 不然看不见结果
	//time.Sleep(time.Second * 5)
	for {
		_ , ok := <- exitChan
		if !ok {
			break
		}
	}
}
