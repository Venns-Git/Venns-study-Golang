package main

import "fmt"

// 声明一种新的数据类型
type myInt int

// 定义一个结构体
type Book struct {
	title string
	auth string
}

func main() {
	var a myInt = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "venns"

	fmt.Printf("%v", book1)
}
