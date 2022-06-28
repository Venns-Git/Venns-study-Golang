package main

import "fmt"

func main() {

	// ========> 第一种声明方式

	// 声明 myMap是一种map类型,key是string,value是string
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("myAap is null map")
	}

	// 使用map前,需要先用make给map分配数据空间
	myMap = make(map[string]string,10)

	myMap["one"] = "java"
	myMap["two"] = "c++"
	myMap["three"] = "golang"

	fmt.Println(myMap)

	// =========> 第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "golang"
	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"one" : "java",
		"two" : "c++",
		"three" : "golang",
	}
	fmt.Println(myMap3)
}
