package main

import "fmt"
//闭包的应用
func main() {
	a := add2(0)

	for i := 0; i < 10; i++ {
		var s int
		s , a = a(i)
		fmt.Println(s)
	}
}

func add() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

type iadd func(int) (int, iadd)

func add2(base int) iadd {
	return func(i int) (int, iadd) {
		return base + i, add2(base + i)
	}
}
