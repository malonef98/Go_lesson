package main

import "fmt"

type Person struct {
	Name string
}

type Student struct {
	Name string
	Ages int
}

func (person Person) test ()  {
	person.Name = "jack"
	fmt.Println("test=:",person.Name)
}

func test2 (p *Person)  {
	fmt.Println("test2:=",p.Name)
}

func main()  {
	//创建结构体的时候直接赋值
	var stu1 = Student{"小明",19}
	sut2 := Student{"小红",18}

	//在创建结构体变量的时候，吧字段名和字段值写在一起
	//这种写法不依赖字段的顺序
	var stu3 = Student{
		Name: "jack",
		Ages: 21,
	}

	stu4 := Student{
		Ages: 17,
		Name: "mary",
	}


	//返回结构体的指针类型
	//返回结构体指针类型
	var stu5  = &Student{"小王",21}
	stu6 := &Student{"小刚",23}

	fmt.Println(stu1,sut2,stu3,stu4,stu5,*stu6)
}
