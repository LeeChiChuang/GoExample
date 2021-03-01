package main

import (
	"fmt"
	"time"
)

func gA(a <-chan int) {
	v := <-a
	fmt.Println("Goroutine A received v:", v)
	return
}

func gB(b <-chan int) {
	v := <-b
	fmt.Println("Goroutine B received v:", v)
	return
}

func main() {
	ch := make(chan int)
	go gA(ch)
	go gB(ch)
	ch <- 3
	time.Sleep(time.Second)
}
