package main

import "fmt"


// 如果类名的首字母大写 表示其他包也能够访问
type Hero struct {
	// 如果类的属性首字母大写 表示该属性可以对外访问
	Name string
	Ad int
	Level int
}

func (this *Hero) Show() {
	fmt.Println("name = ", this.Name)
	fmt.Println("ad = ",this.Ad)
	fmt.Println("level = ",this.Level)
}
func (this *Hero) GetName()  {
	fmt.Println("name = ",this.Name)
}
func (this *Hero) SetName(newName string) {
	this.Name = newName
}
func main() {
	hero := Hero{Name: "Zoe", Ad: 100, Level: 1}
	hero.Show()
	hero.SetName("Timo")
	hero.Show()
}
