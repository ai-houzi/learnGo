package main

import "fmt"


//寻找最长不含有重复字符的子串
func main() {
	fmt.Println(
		lengthOfNoRepetingSubStr("abcd"))
	fmt.Println(
		lengthOfNoRepetingSubStr("abcdasdqwe"))
	fmt.Println(
		lengthOfNoRepetingSubStr("这是我的测试"))
	fmt.Println(
		lengthOfNoRepetingSubStr("一二三三二一"))

}

func lengthOfNoRepetingSubStr(s string) int {
	lastOfOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {

		if last, ok := lastOfOccurred[ch]; ok && last >= start {
			start = last + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOfOccurred[ch] = i
	}
	return maxLength
}
