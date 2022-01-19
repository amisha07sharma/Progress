package main

import (
	"fmt"
)

//go routines

func main() {
	go fmt.Println("Goroutine")
	fmt.Println("Normal")
}
