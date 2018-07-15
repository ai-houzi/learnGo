package main

import (
	"fmt"
	"sync"
)

func main() {
	channelDemo()
}
func channelDemo() {
	var works [10]work

	var wg sync.WaitGroup


	for i := 0; i < 10; i++ {
		works[i] = CreateWork(i, &wg)
	}

	wg.Add(20)

	for i, work := range works {
		work.in <- 'a' + i
	}

	for i, work := range works {
		work.in <- 'A' + i
	}
	wg.Wait()

}

type work struct {
	in chan int
	done func()
}

func CreateWork(id int, wg *sync.WaitGroup) work {
	w := work{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func doWork(id int, w work) {

	for n := range w.in {

		fmt.Printf("第 %d 个work，内容是: %c\n", id, n)
		w.done()
	}

}
