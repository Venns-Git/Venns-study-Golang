package main

import "fmt"

func main() {
	// 固定长度的数组
	var myArray1 [10]int
	myArray2 := [10]int{1,2,3,4,5}

	for i:= 0;i < 10;i++ {
		fmt.Println(myArray1[i])
	}

	for index,value:= range myArray2{
		fmt.Println("index = ", index, "value = ",value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n",myArray1)
}
