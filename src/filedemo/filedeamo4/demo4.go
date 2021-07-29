package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个新文件 并且在文件中写入内容
	// 1 打开文件
	filePath := "/Users/mayifan/Desktop/test1"
	file ,err :=os.OpenFile(filePath,os.O_CREATE | os.O_WRONLY,0666)
	if err != nil{
		fmt.Println("err：",err)
	}

	// 2 写入内容
	str := "hello Gardon\n"
	//使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0 ; i < 5 ; i++ {
		writer.WriteString(str)
	}

	// 及时关闭file句柄
	defer file.Close()

	//writer是带缓存的，内容是先写入缓存的，此时内容还在缓存中，并没有实际写入文件
	//Flush 缓存内容写入文件
	writer.Flush()
}
