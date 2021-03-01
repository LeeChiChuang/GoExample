package main

import (
	"fmt"
	"time"
)

//定时任务
func worker7()  {
	ticker := time.Tick(1*time.Second)
	for {
		select {
		case <- ticker:
			fmt.Println("执行1s任务")
		}
	}
}

func main()  {
	worker7()
}
