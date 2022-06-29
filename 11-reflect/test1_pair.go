package main

import "fmt"

func main() {

	var a string
	a = "hello"
	// pair<type:string,value:"hello">

	var allType interface{}
	allType = a
	//pair<type:string, value:"hello">

	value, _ := allType.(string)
	fmt.Println(value)
}
