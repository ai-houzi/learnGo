package main

import (
	"fmt"
	"learnGo/interface/ayy"
)

type Retriever interface {
	Get(url string) string
}

type Triever struct {
	content string
}

func (t Triever) Get(url string) string {
	return t.content+url
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = ayy.Retriever{}
	//fmt.Println(download(r))

	s := r.(ayy.Retriever)
	fmt.Println(s.TimeOut)
}
