package main

import "fmt"

func main() {
	const (
		b  = 1 << (10 * iota)
		mb
		gb

	)
	fmt.Println(mb)
}
