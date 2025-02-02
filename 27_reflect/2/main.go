
package main

import (
	"fmt"
	"reflect"
)

type Person struct{}

func (p Person) Greet() {
	fmt.Println("Hello!")
}

func main() {
	p := Person{}
	v := reflect.ValueOf(p)
	method := v.MethodByName("Greet")
	method.Call(nil) // Calls Greet()
}
