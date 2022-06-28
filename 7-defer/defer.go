package main

import "fmt"

func deferFunc()  int {
	fmt.Println("defer func call...")
	return 0
}
func returnFunc()  int {
	fmt.Println("return func call...")
	return 0
}
func deferAndReturnFunc() int {
	defer deferFunc()
	return returnFunc()
}
func main() {

	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go1")
	fmt.Println("main::hello go2")

	deferAndReturnFunc()
}
