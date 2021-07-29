package module

type student struct {
	name string
	socre float64
}

//因为student的首字母是小写，属于私有，只能在本包内使用
//工厂模式解决该问题
//返回的是指针，也就是数据本身
func NewStudent(name string,s float64) *student{
	return &student{
		name:  name,
		socre: s,
	}
}

//如果结构体字段是小写，外部包无法直接获取结构体字段，编写方法获取
func  (s student) GetSocre() float64 {
	return s.socre
}