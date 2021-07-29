package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
   //将test文件内容读取到内存  将读取到的内容写入test2

	file1Path := "/Users/mayifan/Desktop/test1"
	file2Path := "/Users/mayifan/Desktop/test"

	content ,err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file err=%v",err)
		return
	}

	err = ioutil.WriteFile(file2Path,content,666)
	if err != nil {
		fmt.Printf("write file err%v",err)
	}
}