package main

import "fmt"

type student struct {

}

func TypeJudge(items... interface{})  {
	for index , x := range items{
		//l类型断言
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值是%v\n", index, x)
		case int:
			fmt.Printf("第%v个参数是 int 类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是 string 类型，值是%v\n", index, x)
		case student:
			fmt.Printf("第%v个参数是 student 类型，值是%v\n", index, x)
		case *student:
			fmt.Printf("第%v个参数是 *student 类型，值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数是 不确定 类型，值是%v\n", index, x)
		}
	}


}


func main() {
	var n1 float32 = 1.1
	var n2 bool = true
	address := "北京"

	TypeJudge(n1,n2,address)
}
