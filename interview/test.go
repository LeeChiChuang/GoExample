package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "hello"
	fmt.Println(strings.Count(str, ""))
	fmt.Println(str[0:1])
	fmt.Println(str)
	fmt.Println(strings.Count(str, ""))
	fmt.Println(str[1])
}
