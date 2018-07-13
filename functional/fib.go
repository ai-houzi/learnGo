package main

import (
	"fmt"
	"strings"
	"io"
	"bufio"
)

//斐波那契数列
func main() {
	f := fibonacci()
	printFile(f)

}

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	s := fmt.Sprintf("%d\n", next)

	if next > 1000 {
		return 0, io.EOF
	}
	return strings.NewReader(s).Read(p)
}

func printFile(reader io.Reader) {

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
