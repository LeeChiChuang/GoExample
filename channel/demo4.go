package main

func main()  {
	//ch := make(chan int)
	var ch chan int
	close(ch)
	//close(ch)
}
// 发生painc的三种情况
//向关闭的channel进行写操作
//关闭一个nil的channel
//重复关闭一个channel
