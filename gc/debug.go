package main

func allocate()  {
	_ = make([]byte, 1 << 20)
}

func main()  {
	for n := 1; n<10000; n++ {
		allocate()
	}
}
