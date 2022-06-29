package main

import "fmt"

// 定义一个接口,本质是一个指针
type AnimaITF interface {
	Sleep()
	GetColor() string
	GetType() string
}


// 具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "Cat"
}
// 具体的类
type Dog struct {
	color string
}
func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GetType() string {
	return "Dog"
}
func main() {
	var animal AnimaITF // 接口的数据类型,父级指针
	animal = &Cat{"red"}
	animal.Sleep()

	animal = &Dog{"green"}
	animal.Sleep()
}
