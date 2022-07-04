package main

import "fmt"

func main() {
	// 定义一个channel
	c := make(chan int)

	go func() {
		fmt.Println("goroutine call..")
		defer fmt.Println("goroutine end")
		c <- 123 // 将123发生给c
	}()
	num := <-c // 从c中接受数据,并赋值给num
	fmt.Println("num = ", num)
	fmt.Println("main goroutine end")
}
