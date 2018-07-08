package main

import "fmt"

func main() {
	var arr1 [5]int 	//定义长度为5的数组
	arr2 := [3]int{1,2,3}	//使用:时需要赋值
	arr3 := [...]int{2,3,4,5}	//可变长度

	fmt.Println(arr1,arr2,arr3)
}
