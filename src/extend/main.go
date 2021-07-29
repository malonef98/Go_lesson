package main

import "fmt"

//学生共有属性
type Student struct {
	Name string
	Age int
	Score int
}

//将Pupil和Graduate共有的方法写到student
func (s *Student) ShowInfo(){
	fmt.Printf("学生姓名%v,学生年龄%v，学生成绩%v\n",s.Name,s.Age,s.Score)
}


type Pupil struct {
	//嵌入匿名结构体
	Student
}

type Graduate struct {
	//嵌入匿名结构体
	Student
}

func (pupil *Pupil) test()  {
	fmt.Println("小学生正在考试......")
}

func (graduate *Graduate) test()  {
	fmt.Println("大学生正在考试......")
}

func main() {
	pupil1 := &Pupil{}
	pupil1.Student.Name = "tom"
	pupil1.Student.Age = 13
	pupil1.Student.Score = 88
	pupil1.Student.ShowInfo()
	pupil1.test()

	graduate1 := &Graduate{}
	graduate1.Student.Name = "marry"
	graduate1.Student.Age = 20
	graduate1.Student.Score = 90
	graduate1.Student.ShowInfo()
	graduate1.test()
}
