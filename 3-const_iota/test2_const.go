package main

import "fmt"
// go可以用const定义枚举类型
const (
	// 可以在const()添加一个关键字iota,每行的iota都会累加1,从0开始
	BEIJING  = iota // iota = 0
	SHANGHAI		// iota = 1
	SHENZHEN 		// iota = 2
)
const (
	a,b = iota + 1,iota + 2 // iota = 0, a = 1, b = 2
	c,d						// iota = 1, c = 2, d = 3
	e,f						// iota = 2, e = 3, f = 4
	g,h = iota * 2,iota * 3	// iota = 3, g = 6, h = 9
)

func main() {

	// 常量(只读类型)
	const length int = 100
	fmt.Println("length = ",length)


	fmt.Println("BEIJING = ",BEIJING)
	fmt.Println("SHANGHAI = ",SHANGHAI)
	fmt.Println("SHENZHEN = ",SHENZHEN)

	fmt.Println("a = ",a,"b = ",b)
	fmt.Println("c = ",c,"d = ",d)
	fmt.Println("e = ",e,"f = ",f)
	fmt.Println("g = ",g,"h = ",h)

}
