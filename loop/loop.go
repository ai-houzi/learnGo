package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(0),
	)

	printFile("branch/abc.txt")
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result += strconv.Itoa(lsb)
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("读取出错")
	}
	sacnner := bufio.NewScanner(file)

	for sacnner.Scan() {
		fmt.Println(sacnner.Text())
	}
}
