package main

import "fmt"

type Person struct {
	Name string
	Age int
	Edu string
	Address string
}

//给Person类型绑定方法
func(person Person) getName() {
	fmt.Println("Name:",person.Name)
}

func (person Person) speak() {
	fmt.Printf("%v是个好人",person.Name)
}

func (Person Person) jisuan(n int) {
	i := 0
	j := 1
	for j <= n {
		i = i + j
		j++
	}
	fmt.Println(i)
}

func main() {
	var p Person
	var b Person
	p.Name = "Tom"
	b.Name = "jerry"
	b.speak()
	b.jisuan(10)
}
