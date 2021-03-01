package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
翻转字符串
问题描述

请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。

给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。
 */

func reverString(s string) (string, bool) {
	str := []rune(s)
	l := len(str)
	if l > 5000 {
		return s, false
	}
	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	return string(str), true
}

func string2Bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func s2b(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len

	return b
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func main()  {
	str := "hello"
	b1 := s2b(str)
	/*
	l := len(h)
	for i := 0; i < l/2; i++ {
		h[i], h[l-1-i] = str[l-1-i], str[i]
	}
	 */
	fmt.Println(b1)
	b1[0] = 'a'

	b2 := StringToBytes(str)
	fmt.Println(b2)
	b2[0] = 'a'
	//fmt.Println(string(h))
}
