package main

import "fmt"

// 单个返回值
func func1(a string,b int) int{
	fmt.Println("a = ", a, "b = ",b)
	c := 100
	return c
}
// 多个返回值,匿名
func func2(a string,b int) (int ,int){
	fmt.Println("a = ", a, "b = ",b)
	return 111,222
}
// 多个返回值,有形参名称
func func3(a string,b int) (r1 int,r2 int) {
	fmt.Println("a = ",a,"b = ",b)
	// r1,r2相当于函数的局部变量,默认值为0
	fmt.Println("r1 = ", r1, "r2 = ", r2)
	r1,r2 = 111,222
	return
}
func func4(a string,b int) (r1, r2 int) {
	fmt.Println("a = ",a,"b = ",b)
	r1,r2 = 111,222
	return
}

func main() {
	c := func1("hello",123)
	fmt.Println("c = ",c)

	ret1,ret2 := func2("abc",123)
	fmt.Println("ret1 = ",ret1,"ret2 = ret2 = ",ret2)

	ret1,ret2 = func3("abc",123)
	fmt.Println("ret1 = ",ret1,"ret2 = ret2 = ",ret2)

	ret1,ret2 = func4("abc",123)
	fmt.Println("ret1 = ",ret1,"ret2 = ret2 = ",ret2)
}
