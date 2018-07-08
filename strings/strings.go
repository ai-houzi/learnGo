package main

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	s := "我爱houzi！"

	bytes := []byte(s)

	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", r)
	}
	fmt.Println()
	for r, ch := range []rune(s) {
		fmt.Printf("(%d,%c)", r, ch)
	}
	fmt.Println()
}
