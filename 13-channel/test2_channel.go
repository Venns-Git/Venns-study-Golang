package main

import "fmt"

func main() {
	c := make(chan int, 3) // 带有缓冲的channel

	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("goroutine end")

		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("goroutine call,发送的元素 = ", i, "len(c) = ", len(c), "cap(c) = ", cap(c))
		}
	}()

	for i := 0; i < 3; i++ {
		num := <-c // 从c中接收数据,并赋值给num
		fmt.Println("num = ", num)
	}

	fmt.Println("main end")
}
