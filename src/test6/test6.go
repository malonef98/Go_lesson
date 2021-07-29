package main

import "fmt"

type Circle struct {
	redius float64
}

func (c Circle) area () float64  {
	c.redius = 8.0
	return 3.14 * c.redius * c.redius
}

//通常为了提高效率，我们的方法和结构体指针进行绑定
func (c *Circle) area2 () float64  {
	c.redius = 9.0
	return 3.14 * c.redius * c.redius
}

type integer int
func (i integer) print() {
	fmt.Println("i=",i)
}

func main(){
	 var i integer
	 i = 100
	 i.print()
}
