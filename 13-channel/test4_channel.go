package main

import "fmt"

/*
	select 可以监控多个channel
*/
func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			// 如果c可写,则该case就会进来
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			// 如果c中没有数据,则会阻塞
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacii(c, quit)

}
