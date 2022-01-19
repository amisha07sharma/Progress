package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Anonymous goroutine")
		}
	}()

	go fmt.Println("Goroutine")
	time.Sleep(time.Second)
	fmt.Println("Normal")

}
