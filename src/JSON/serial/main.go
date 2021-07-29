package main

import (
"encoding/json"
"fmt"
)

type Monster struct {
	//使用json的tag  改变序列化之后的字段名字
	Name string     `json:"monster_name"`
	Age int			`json:"monster_age"`
	Birthday string
	Sal string
	Skill string
}

func testSsruct()  {
	monster := Monster{
		Name: "牛魔王",
		Age: 500,
		Birthday: "2020-1-1",
		Sal: "8000.0",
		Skill: "牛魔拳",
	}

	data,err := json.Marshal(&monster)
	if err !=nil {
		fmt.Println("序列化失败！！！")
	}
	fmt.Printf("%v",string(data))
}

//map的序列化
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	data, err := json.Marshal(&a)
	if err != nil {
		fmt.Println("序列化失败！！！")
	}
	fmt.Printf("%v",string(data))
}


//切片的序列化
func testSlice()  {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	var m2 map[string]interface{}
	m1 = make(map[string]interface{})
	m2 = make(map[string]interface{})

	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = "beijing"

	m2["name"] = "tom"
	m2["age"] = 17
	m2["address"] = "shanghai"

	slice = append(slice,m1)
	slice = append(slice,m2)

	data, err := json.Marshal(&slice)
	if err != nil {
		fmt.Println("序列化失败！！！")
	}
	fmt.Printf(string(data))
}

func testFlaot()  {
	num := 234.5
	data, err := json.Marshal(&num)
	if err != nil {
		fmt.Println("序列化失败！！！")
	}
	fmt.Printf("%v",string(data))

}

func main() {
     testSlice()
}
