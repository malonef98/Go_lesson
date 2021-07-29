package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开已经存在的内容，覆盖文件内容
	// 1 打开文件
	filePath := "/Users/mayifan/Desktop/test1"
	//os.O_RDWR 读写的方式打开
	file ,err :=os.OpenFile(filePath,os.O_APPEND | os.O_RDWR,0666)
	if err != nil{
		fmt.Println("err：",err)
	}

	//读取文件原内容
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Println(str)
	}

	// 写入追加内容
	str := "Hello BeiJin!!!\n"
	//使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0 ; i < 10 ; i++ {
		writer.WriteString(str)
	}

	// 及时关闭file句柄
	defer file.Close()

	//writer是带缓存的，内容是先写入缓存的，此时内容还在缓存中，并没有实际写入文件
	//Flush 缓存内容写入文件
	writer.Flush()
}
