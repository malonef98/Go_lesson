package main

import (
	"fmt"
	"strconv"
)

func main(){
	//统计字符串的长度
	str := "golang你好"     //UTF-8字符集 一个汉字3字节
	fmt.Println(len(str)) //12字节

	//字符串遍历
	// 方式1 利用键值循环遍历：for-range
	for i ,value := range str{
		fmt.Printf("%d,%c\n",i,value)
	}

	// 方式2 利用切片
	r := []rune(str)
	for i := 0; i <len(r); i++ {
		fmt.Printf("%c",r[i])
	}

	n := strconv.Itoa(66)
	fmt.Printf(n)
}



