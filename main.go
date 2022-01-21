package main

import (
	"fmt"
	"time"
)

type Object struct {
}

func printText() {
	for i := 0; i < 5; i++ {
		fmt.Println("Anonymous goroutine")
	}
}

func ping(ch chan Object, i Object) {
	ch <- i
	fmt.Println("ping")
}

func pong(ch chan Object) {
	<-ch
	fmt.Println("pong")
}

func main() {

	ch := make(chan Object)

	for i := 0; i < 5; i++ {

		go pong(ch)
		time.Sleep(time.Second)

		go ping(ch, Object{})
		time.Sleep(time.Second * 2)
	}

	close(ch)
	go printText()
	time.Sleep(time.Second)

	fmt.Println("Normal")
}

// 5 rounds
// ping pong
// once ping then pong
// what is functional programmimg, how functional programming can be achieved using golang.
// currying
// dining philosophers problem
