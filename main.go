package main

import (
	"fmt"
	"time"
)

func printText() {
	for i := 0; i < 5; i++ {
		fmt.Println("Anonymous goroutine")
	}
}

func main() {

	// ping = recieving
	// pong = sending
	ch := make(chan int32, 8)

	time.Sleep(time.Second * 5)

	ch <- 3
	go fmt.Println("ping")
	time.Sleep(time.Second * 5)

	i := <-ch

	go fmt.Println(i, "pong")

	go printText()

	time.Sleep(time.Second * 5)

	fmt.Println("Normal")
}

// 5 rounds
// ping pong
// once ping then pong
// what is functional programmimg, how gunctional programming can be achieved using golang.
// currying
// dining philosophers problem
