package main

import (
	"fmt"
	"sync"
)

/*
交替打印数字和字母
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

func main() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

		i := 0
		for {
			select {
			case <-letter:
				if i >= len(str) {
					wait.Done()
					return
				}

				fmt.Print(str[i: i+1])
				i++
				if i >= len(str) {
					i = 0
				}
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}

		}
	}(&wait)
	number <- true
	wait.Wait()
}
