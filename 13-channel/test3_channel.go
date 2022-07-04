package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// close 可以关闭一个channel
		close(c)
	}()

	/*
		1.采用死循环
		for {
			// ok表示channel是否关闭,true为未关闭
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	*/
	// 2. 使用range
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main end")
}
