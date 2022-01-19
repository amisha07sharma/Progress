package main

import (
	"fmt"
	"time"
)

//go routines

func main() {
	go fmt.Println("Goroutine")
	time.Sleep(time.Second * 2)
	fmt.Println("Normal")
}
