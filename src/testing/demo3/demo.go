package demo3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

func (m *Monster) Store()  bool {

	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("序列化失败！！！")
	}
	fmt.Printf("%v",string(data))

	//保存到文件
	// 1 打开文件
	filePath := "/Users/mayifan/Desktop/monster.log"

	// 2 写入内容
	//使用带缓存的 *Writer
	err = ioutil.WriteFile(filePath,data,666)
	if err != nil {
		fmt.Println("write file err =",err)
		return false
	}
	return true
}

func (m *Monster) ReStore() bool {
	filePath := "/Users/mayifan/Desktop/monster.log"

	data , err := ioutil.ReadFile(filePath)
	if err !=nil {
		fmt.Println("read file err =",err)
		return false
	}

	//使用读取到的data []byte 反序列化
	var monster Monster
	err = json.Unmarshal(data,&monster)
	if err != nil {
		fmt.Println("unmarshal err = ",err)
		return false
	}
	return true
}