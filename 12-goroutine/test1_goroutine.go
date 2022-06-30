package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	for i := 0; i < 10; i++ {
		fmt.Println("new Goroutine: i = ", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	// 创建一个go程 去执行newTask() 流程
	go newTask()

	for i := 0; i < 10; i++ {
		fmt.Println("main Goroutine:i = ", i)
		time.Sleep(1 * time.Second)
	}
}
