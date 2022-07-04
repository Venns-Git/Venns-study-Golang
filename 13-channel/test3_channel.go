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

	for {
		// ok表示channel是否关闭,true为未关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main end")
}
