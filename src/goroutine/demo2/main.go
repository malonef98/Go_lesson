package main

import (
	"fmt"
	"sync"
)

var (
	myMap = make(map[int]int,10)

	//声明一个全局的互斥锁
	//sync 是包 低水平程序线程使用
	//Mutex 互斥
	lock sync.Mutex
)

//阶乘运算
func test(n int)  {
	res := 1
	for i := 1 ; i <= n ; i++ {
		res *=i
	}

	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main() {
	//开启200个协程来进行 200个阶乘计算
	for i := 1; i <= 20 ; i++ {
		go test(i)
	}

	lock.Lock()
	//输出结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n" ,i,v)
	}
	lock.Unlock()
}
