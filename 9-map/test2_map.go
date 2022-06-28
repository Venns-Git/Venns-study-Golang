package main

import "fmt"

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["china"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	for key, val := range cityMap {
		fmt.Println("key = ", key, "value = ", val)
	}

	// 删除
	delete(cityMap,"USA")

	fmt.Println("=============")

	// 遍历
	for key, val := range cityMap {
		fmt.Println("key = ", key, "value = ", val)
	}
}
