package main

import "fmt"

//寻找最长不含有重复字符的子串
func main() {
	fmt.Println(
		LengthOfNoRepetingSubStr("abcd"))
	fmt.Println(
		LengthOfNoRepetingSubStr("abcdasdqwe"))
	fmt.Println(
		LengthOfNoRepetingSubStr("这是我的测试"))
	fmt.Println(
		LengthOfNoRepetingSubStr("一二三三二一"))
	fmt.Println(
		LengthOfNoRepetingSubStr("黑灰化肥灰会挥发发灰黑讳为黑灰花会飞"))

}

var lastOfOccurred = make([]int, 0xffff)

func LengthOfNoRepetingSubStr(s string) int {
	//lastOfOccurred := make(map[rune]int)
	 //lastOfOccurred := make([]int, 0xffff)

	for i := range lastOfOccurred {
		lastOfOccurred[i] = -1
	}
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {

		if last := lastOfOccurred[ch]; last != -1 && last >= start {
			start = last + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOfOccurred[ch] = i
	}
	return maxLength
}
