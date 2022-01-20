package main

import (
	"GO/math"
	"fmt"
)

func printText() {
	for i := 0; i < 5; i++ {
		fmt.Println("Anonymous goroutine")
	}
}

func ping(ch chan int, i int) {
	ch <- i
	fmt.Println("ping", i)
}

func pong(ch chan int) {
	i := <-ch
	fmt.Println("pong", i)
}

func main() {

	// ch := make(chan int)

	// for i := 0; i < 5; i++ {

	// 	go pong(ch)
	// 	time.Sleep(time.Second)

	// 	go ping(ch, i)
	// 	time.Sleep(time.Second * 2)
	// }

	// close(ch)
	// go printText()
	// time.Sleep(time.Second)

	// fmt.Println("Normal")

	math.Add(2, 3)
	math.Subtract(9, 2)
}

// 5 rounds
// ping pong
// once ping then pong
// what is functional programmimg, how functional programming can be achieved using golang.
// currying
// dining philosophers problem
