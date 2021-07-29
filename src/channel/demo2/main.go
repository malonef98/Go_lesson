package main

import "fmt"

type Cat struct {
	Name string
	Age int
}

func main() {
	allChan := make(chan interface{},3)

	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"小花猫",4}
	allChan <- cat

	//队列 先进先出 要先推出前两个变量
	<- allChan
	<- allChan

	//获取第三个变量
	newCat := <-allChan

	fmt.Printf("newCat=%T，newCat=%v",newCat,newCat)

	//类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v",a.Name)
	//fmt.Printf("newCat.Name=%v",newCat.Name)     错误，编译器任务newCat为借口类型
}
