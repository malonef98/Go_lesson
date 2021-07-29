package main

import (
	"fmt"
	"strconv"
	"time"
)

//编写一个函数 每隔一秒输出hello world
func test() {
	for i := 1 ; i <= 10 ; i++ {
		fmt.Println("hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}


func main() {

	//开启一个协程
	go test()

	for i := 1; i <= 10 ; i ++ {
		fmt.Println("mian()  hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
