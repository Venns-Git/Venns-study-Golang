package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"2-4 length"`
	Sex  string `info:"sex" doc:"type is female or male"`
}

func findTag(arg interface{}, tagName string) {
	t := reflect.TypeOf(arg).Elem()
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get(tagName)
		fmt.Println(tagName, ": ", tag)
	}
}

func main() {
	var re resume
	findTag(&re, "doc")
}
