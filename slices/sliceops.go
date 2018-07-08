package main

import "fmt"

func main() {
	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s,i*2)
	}

	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)


	fmt.Println("slice 拷贝")

	copy(s2,s1)
	fmt.Println(s2)

	fmt.Println("slice 删除8元素")
	s2 = append(s2[:3],s2[4:]...)
	fmt.Println(s2)
	printSlice(s2)

	fmt.Println("删除头元素")

	front := s2[0]
	s2 = s2[1:]

	fmt.Println("删除尾元素")

	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(front,tail)
	printSlice(s2)

}

func printSlice(s []int)  {
	fmt.Printf("len = %d,cap = %d\n",len(s),cap(s))
}