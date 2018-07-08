package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	const filename = "branch/abc.txt"
	if constens, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", constens)
	}
}
//switch后面可以没有表达式
func grade(score int) string {
	g := ""
	switch {
	case score<0||score>100:
		panic(
			fmt.Sprintf("Wrong score: %d",score))
	case score<60:
		g = "F"
	case score<80:
		g = "C"
	case score<90:
		g = "B"
	case score<=100:
		g = "A"
	}
	return g
}
