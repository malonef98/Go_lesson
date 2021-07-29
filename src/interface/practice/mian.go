package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1 声明一个Hero结构体
type Hero struct {
	Name string
	Age int
	Score int
}

// 2 声明结Hero构体切片类型
type HeroSlice []Hero

// 3 实现排序interface接口
// 实现Len函数
func (hs HeroSlice) Len()  int {
	 return len(hs)
}

// Less 方法就是决定你使用什么标准进行排序
// 按照Hero的成绩从小到大
func (hs HeroSlice) Less(i,j int) bool {
	return hs[i].Score > hs[j].Score
}

// Swap 进行变换
func (hs HeroSlice) Swap(i,j int)  {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func main() {
	var heroes HeroSlice
	for i := 0 ; i < 10 ; i++ {
		hero := Hero{
			Name: fmt.Sprint("英雄",rand.Intn(100)),
			Age: rand.Intn(100),
			Score: rand.Intn(100),
		}
		//将英雄加入切片
		heroes = append(heroes,hero)
	}


	i := 20
	j := 10
	//交换的快速写法
	i,j = j,i

	fmt.Println(i,j)
	//排序前的顺序
	fmt.Println(heroes)
	// 结构体已经实现对应的sort接口方法
	sort.Sort(heroes)
	//排序后的顺序
	fmt.Println(heroes)

}
