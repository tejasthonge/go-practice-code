package main

import (
	"fmt"
	"reflect"
	"time"
)

/*
The reflect package in Go (reflect) is used for runtime reflection,
allowing you to inspect and manipulate types, values, and structures dynamically.
It is useful for handling dynamic types, inspecting struct fields,
and working with interfaces.
*/
func main() {
	name := "Tejas"
	comName := "qbitrom"
	age := 23
	fmt.Println(reflect.TypeOf(name)) // it return the reflect.Type ,string
	founded := time.Now()
	fmt.Println(reflect.TypeOf(comName)) //string
	fmt.Println(founded)                 //2025-02-02 19:10:22.039629 +0530 IST m=+0.004603001

	// var i interface{}
	// i.(type) // errer we cant not user .(type outside the switch)

	refVal := reflect.ValueOf(name) //return the  reflect.Value
	fmt.Println(refVal)
	refV := reflect.ValueOf(age).Interface().(string)

	fmt.Println(refV)

}
