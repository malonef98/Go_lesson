package main

import "fmt"

type MethodUtils struct {
}

func (M *MethodUtils) print ()  {
	for i := 0 ; i < 8 ; i++ {
		for j := 0 ; j < 10 ; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (M *MethodUtils) print1 (n int, m int) {
	for i := 0 ; i < n ; i++ {
		for j := 0 ; j < m ; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (M *MethodUtils) mianji (len int, width int) {
	fmt.Printf("面积为：%v",len * width)
}

func (M *MethodUtils) JudgeNum (num int) {
	if num % 2 == 0 {
		print("偶数")
	}else {
		print("奇数")
	}
}

func (M *MethodUtils) Jisuanji (num1 int, num2 int, op string) float64{
	switch op {
	case "-":
		return (float64(num1 - num2))
	case "+":
		return (float64(num1 + num2))
	case "*":
		return (float64(num1 * num2))
	case "/":
		return (float64(num1 / num2))
	default:
		return 0
	}
}

func main() {
	var m MethodUtils
	fmt.Println(m.Jisuanji(8,2,"/"))
}
