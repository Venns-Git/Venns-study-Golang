package main

import "fmt"
// 声明全局变量,方法1,2,3都可以,方法4都只能在函数体内使用
var gA int = 100
var gB = 200

// 变量4种声明方式
func main() {
	// 1. 声明一个变量,默认值是0
	var a int
	fmt.Println("a = ",a)
	fmt.Printf("type of a = %T\n",a)

	// 2. 声明一个变量,初始化一个值
	var b int = 100
	fmt.Println("b = ",b)
	fmt.Printf("type of b = %T\n",b)

	var bb string = "hello world"
	fmt.Println("bb = ",bb)
	fmt.Printf("type of bb = %T\n",bb)

	// 3. 在初始化时,可以省略数据类型,自动匹配
	var c = 100
	fmt.Println("c = ",c)
	fmt.Printf("type of c = %T\n",c)

	// 4. 省略var关键字,直接自动匹配(常用)
	d := 100
	fmt.Println("d = ",d)
	fmt.Printf("type of d = %T\n",d)

	e := "abcd"
	fmt.Println("e = ",e)
	fmt.Printf("type of e = %T\n",e)

	f := 3.14
	fmt.Println("f = ",f)
	fmt.Printf("type of f = %T\n",f)

	// ====
	fmt.Println("gA = ",gA,",gB = ",gB)

	// 声明多个变量
	var xx,yy int = 100,200
	fmt.Println("xx = ",xx,",yy = ",yy)

	var kk,ll = 100,"hello"
	fmt.Println("kk = ",kk,",ll = ",ll)

	// 多行的变量声明
	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println("vv = ",vv,",jj = ",jj)
}
