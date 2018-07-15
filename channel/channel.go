package main

import (
	"fmt"
	"time"
)

func main() {
	//channelDemo()

	//bufferdChannel()

	chanmelClose()
}
func channelDemo() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = CreateWork(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func CreateWork(id int) chan<- int {
	c := make(chan int)
	go work(id, c)
	return c
}

func work(id int, c chan int) {

	for n := range c {

		fmt.Printf("第 %d 个work，内容是: %c\n", id, n)
	}
}
func bufferdChannel() {
	c := make(chan int, 3)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func chanmelClose() {
	c := make(chan int)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
