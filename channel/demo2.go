package main

import "fmt"

// 单向通道
//单向只写
func counter(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func squarer(out chan<- int, in <-chan int) {
	for {
		v, ok := <-in
		if !ok {
			break
		}
		out <- v * v
	}
	close(out)
}

//单向只读
func printer(c <-chan int) {
	for i := range c {
		fmt.Println("i:", i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
