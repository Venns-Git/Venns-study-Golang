package main

import "fmt"

// interface{} 是万能数据类型

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// interface{} 如何区分此时引用的底层数据类型到底是什么,可以使用"类型断言"机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)
	}
}

type Message struct {
	auth string
}

func main() {
	m := Message{"Golang"}
	myFunc(m)
	myFunc(100)
	myFunc("hello")
}
