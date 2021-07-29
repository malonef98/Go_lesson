package main

import (
	"errors"
	"fmt"
)

func main() {
    err := test()
    if err != nil{
    	fmt.Println("自定义错误",err)
	}
    fmt.Println("上面除法函数执行成功")
}

func test() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		//抛出自定义错误
		return errors.New("除数不能为零")
	} else {
		//无错误，继续执行
		result := num1 / num2
		fmt.Println(result)
		//无错误，返回空值
		return nil
	}
}