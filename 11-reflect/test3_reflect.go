package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "venns", 18}
	DoFieldAndMethod(user)
}

func DoFieldAndMethod(arg interface{}) {
	// 获取arg的type
	argType := reflect.TypeOf(arg)
	fmt.Println("argType is ", argType.Name())

	// 获取arg的value
	argVal := reflect.ValueOf(arg)
	fmt.Println("argValue is", argVal)

	// 获取arg type里面的字段
	// 1. 通过type得到field, 进行遍历
	// 2. 得到每个field
	// 3. 通过field得到对应的value
	for i := 0; i < argType.NumField(); i++ {
		field := argType.Field(i)
		value := argVal.Field(i).Interface()
		fmt.Printf("%s : %v\n", field.Name, value)
	}

	// 获取arg type里面的方法
	for i := 0; i < argType.NumMethod(); i++ {
		method := argType.Method(i)
		fmt.Printf("%s : %v\n", method.Name, method.Type)
	}
}
