package main

import (
	"fmt"
	"os"
)

func main() {

	//打开文件
	// 叫法不同，概念相同
	// 1 file 叫file对象
	// 2 file 叫文件指针
	// 3 file 叫文件句柄
	file , err := os.Open("/Users/mayifan/Desktop/test.txt")
	if err != nil {
		fmt.Println("open file err = ",err)
	}

	fmt.Printf("file=%v",file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=",err)
	}
}
