package main

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

//如果实现A接口，就必须需要实现B，C接口的方法
type AInterface interface {
	BInterface
	CInterface
	test03()
}

type Stu struct {

}

func (s Stu) test01()  {

}

func (s Stu) test02()  {

}

func (s Stu) test03()  {


}
func main() {
	var stu  Stu
	var a AInterface = stu
	 a.test01()
}
