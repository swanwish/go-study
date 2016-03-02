package main

import (
	"fmt"
	"reflect"
)

func main() {
	testReflect()
}

func testReflect() {
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	t := v.Type()
	fmt.Println(t.String())
}
