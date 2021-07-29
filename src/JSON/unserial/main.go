package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name string
	Age int
	Birthday string
	Sal string
	Skill string
}

func unserialStruct()  {
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2020-1-1\",\"Sal\":\"8000.0\",\"Skill\":\"牛魔拳\"}\n"
	var monster Monster
	err := json.Unmarshal([]byte(str),&monster)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Printf("%v",monster)
}

func unserialMap(){
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"

	//map的结构和序列化之前要一致
	var m map[string]interface{}

	//反序列化无需再make map
	err := json.Unmarshal([]byte(str),&m)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Printf("%v",m)

}

func unerialSlice(){
	str := "[{\"address\":\"beijing\",\"age\":7,\"name\":\"jack\"},{\"address\":\"shanghai\",\"age\":17,\"name\":\"tom\"}]\n"

	//切片类型要和序列化之前一致
	var slice []map[string]interface{}

	//反序列化无需再make map
	err := json.Unmarshal([]byte(str),&slice)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Printf("%v",slice)
}

func main() {
	unerialSlice()
}
