package main
import "fmt"

func main()  {
	//进行类型转换
	var n1 int = 100
	//var n2 float32 = n1 自动转换不好使，必须显示转换
	fmt.Println(n1)

	var n2 float32 = float32(n1)
	fmt.Println(n2)
	// n1 的类型没有改变，只是将它的值改为float32类型

	//可以正常编译输出，但是数据会溢出，精度的丧失
	var n3 int64 = 888888
	var n4 int8  = int8(n3)
	fmt.Println(n4)  //56

	var n5 int32 = 12
	//必须将n5转成int64，与等号两边类型一致
	var n6 int64 = int64(n5) + 30
	fmt.Println(n5)
	fmt.Println(n6)

	var n7 int64 = 15
	var n8 int8 = int8(n7) + 127   //编译通过，但是结果可能会溢出
	//var n9 int8 = int8(n7) + 128   //编译不会通过，已经超过类型范围
	fmt.Println(n8)
	//fmt.Println(n9)
}
