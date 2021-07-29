package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file , err := os.Open("/Users/mayifan/Desktop/test")
	if err != nil {
		fmt.Println("open file err = ",err)
	}

	fmt.Printf("file=%v\n",file)

	//函数结束时候及时关闭file 否则会内存泄漏
	defer file.Close()

	//创建一个Reader  带缓存区
	reader := bufio.NewReader(file)
	for  {
		//读到一个换行符就结束一次
		str , err := reader.ReadString('\n')
		//io.EOF是读取到文件末尾
		if err == io.EOF {
			break
		}
		//输出内容
		fmt.Print(str)
	}

	fmt.Println("文件读取结束")
}