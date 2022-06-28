package main // main函数必须在main下

import (
	"fmt"
	"time"
) // 多个包导入方式

func main() { // 函数左括号一定和函数名在同一行
	fmt.Println("hello Go!")
	time.Sleep(1 * time.Second)
}
