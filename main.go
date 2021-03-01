package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var a float64
	a = 3.14
	b := uint64(a)
	c := *(*uint64)(unsafe.Pointer(&a))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
