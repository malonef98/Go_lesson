package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "/Users/mayifan/Desktop/test"
	//使用ioutil.ReadFile一次性读取文件
	comtent, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file err:=%v",err)
	}

	//把读取的文件内容显示到终端
	fmt.Printf("%v",string(comtent))  //[]byte
	//因为 没有显式的Open文件，不需要显示的Close文件
	// Tips 适合小文件的读取 大文件读取效率有限


}
