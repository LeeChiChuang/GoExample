package main

import (
	"sync"
)

func cat(c <-chan bool, d chan<- bool)  {
	for {
		select {
		case <-c:
			println("cat...")
			d<-true
		default:
			break
		}
	}
}

func dog(d <-chan bool, f chan<- bool)  {
	for {
		select {
		case <-d:
			println("dog...")
			f<-true
		default:
			break
		}
	}
}

func fish(f <-chan bool, c chan<- bool, wg *sync.WaitGroup)  {
	i := 0
	for {
		select {
		case <-f:
			println("fish...")
			i++
			if i<100 {
				c<-true
			} else {
				wg.Done()
			}
		default:
			break
		}
	}
}

func main()  {
	wg := sync.WaitGroup{}
	wg.Add(1)
	c := make(chan bool)
	d := make(chan bool)
	f := make(chan bool)
	go cat(c, d)
	go dog(d, f)
	go fish(f, c, &wg)
	c<-true
	wg.Wait()
}
