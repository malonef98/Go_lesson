package main

import (
	"fmt"
)

func PutNum(intChan chan int) {

	for i := 1; i <= 8000; i++ {
		intChan <- i
	}

	//关闭intChan
	close(intChan)
}

func primeNun(intChan chan int, primeChan chan int, exitChan chan bool) {
	//使用for循环
	var num int
	var flag bool
	var ok bool
	for {
		flag = true
		num, ok = <-intChan
		//fmt.Println("\nnum=",num)
		if !ok {
			//管道为空
			break
		}
		//判断num是否为素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				//说明num不是素数
				//fmt.Println("\n不是是素数")
				flag = false
				break
			}
		}
		if flag {
			//是素数  将这个数放入pirmeChan
			//fmt.Printf("\n%v是素数放入",num)
			primeChan <- num
		}
	}

	fmt.Println("\n有一个primeNum 协程取不到数据  退出")
	//这时候还不能关闭管道
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	//标识退出的管道
	exitChan := make(chan bool, 4)

	//开启一个协程  向intChan放入数据 1-8000个数
	go PutNum(intChan)

	//开启4个协程 从intChan中取出数据  并且判断是否为素数  如果是就放入到primeChan中
	for i := 0; i < 4; i++ {
		go primeNun(intChan, primeChan, exitChan)
	}

	//这主线程要进行处理  判断四个协程是否都完成
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//当我们从exitChan取出四个结果 就那个放心的关闭primeChan管道
		close(primeChan)
	}()

	//遍历primeChan 取出结果
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("\n素数为=%v", res)
	}

	fmt.Println("main线程退出")
}
