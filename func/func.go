package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func main() {
	fmt.Println(apply(pow, 3, 4))
	//内部匿名函数
	fmt.Println(
		apply(func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 2),
	)
	fmt.Println(sum(1, 2, 3, 4))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}

func eval(a, b int, op string) (int, error) {
	switch op {

	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("方法错误:" + op)
	}
}

//带余数除法
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

//函数式方法
func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()

	opName := runtime.FuncForPC(pointer).Name()

	fmt.Printf("calling fucntion %s with args (%d ,%d)", opName, a, b)

	return op(a, b)

}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数
func sum(num ...int) int {
	s := 0
	for i := range num {
		s += num[i]
	}
	return s
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
