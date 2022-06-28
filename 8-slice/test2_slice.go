package main

import "fmt"

func printArray(arr []int)  {
	// 引用传递
	for _,val := range arr{
		fmt.Println("value = ",val)
	}
	arr[0] = 100
}

func main() {
	myArray := []int{1,2,3,4} // 动态数组 切片 slice

	fmt.Printf("myArray type is %T\n",myArray)

	printArray(myArray)

	fmt.Println("-------------------")
	for _,val := range myArray{
		fmt.Println("value = ",val)
	}
}
