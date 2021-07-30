package main

import (
	"fmt"
	"time"
)

var (
	chanNun = make(chan int, 2000)
	chanRes = make(chan int, 2000)
)

func PutNum(n int) {
	for i := 1; i <= n; i++ {
		chanNun <- i
	}
	close(chanNun)
}

func Cal() {
	for {
		v, ok := <-chanNun
		if !ok {
			break
		}
		var res int
		for i := 0; i < v; i++ {
			res = res + i
		}
		chanRes <- res
	}
}

func main() {
	go PutNum(2000)

	for i := 0; i < 8; i++ {
		go Cal()
	}

	time.Sleep(time.Second)
	_, ok := <-chanNun
	if !ok {
		close(chanRes)
		for v := range chanRes {
			fmt.Println(v)
		}
	}
}
