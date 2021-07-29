package main

import (
	"fmt"
)

func main() {
	arr := []int{33,12,34,20,55,35,76,6}
	bubble(&arr)
	fmt.Println(arr)
}

func sort1(arr []int)  []int {
	for i := 0 ; i < len(arr) - 1; i++ {
		for j := i; j < len(arr) ; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return  arr
}

func bubble(arr *[]int){
	for i := 0 ; i < len(*arr) - 1; i++ {
		for j := i; j < len(*arr) ; j++ {
			if (*arr)[i] > (*arr)[j] {
				temp := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = temp
			}
		}
	}
}
