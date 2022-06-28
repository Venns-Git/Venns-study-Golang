package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5) // 类型, 长度, 容量

	fmt.Printf("len = %d, cap = %d,slice = %v\n",len(numbers), cap(numbers), numbers)

	// 表示向numbers追加一个元素1
	numbers = append(numbers,1)
	fmt.Printf("len = %d, cap = %d,slice = %v\n",len(numbers), cap(numbers), numbers)

	numbers = append(numbers,2)
	fmt.Printf("len = %d, cap = %d,slice = %v\n",len(numbers), cap(numbers), numbers)

	// 向一个容量已满的slice追加元素
	numbers = append(numbers,3)
	fmt.Printf("len = %d, cap = %d,slice = %v\n",len(numbers), cap(numbers), numbers)
}