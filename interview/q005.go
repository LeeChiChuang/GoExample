package main

import (
	"fmt"
	"sync"
)

/*
在 golang 协程和channel配合使用
写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，
另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。
*/

func main()  {
	out := make(chan int, 5)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i:=0; i<5; i++ {
			out <- i
		}
		close(out)
	}()

	go func() {
		defer wg.Done()
		for v := range out{
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
