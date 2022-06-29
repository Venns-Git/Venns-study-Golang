package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human // SuperMan类继承类Human类的方法
	level int
}

// 重定义父类的方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}
func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}
func main() {
	h := Human{"venns", "female"}

	h.Eat()
	h.Walk()

	// 定义一个子类对象
	s := SuperMan{Human{"superMan", "female"}, 1}
	s.Walk() // 父类的方法
	s.Eat()  // 子类的方法
	s.Fly()  // 子类的方法
	s.Print()

}
