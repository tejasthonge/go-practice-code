package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	//Join ->string
	strSli := []string{"Tejas", "Thonge", "hidu", "Maratha"}
	fmt.Println(reflect.TypeOf(strSli), strSli) //[]string [Tejas Thonge hidu Maratha]
	joinStr := strings.Join(strSli, ",")

	fmt.Println(reflect.TypeOf(joinStr), joinStr) //string Tejas,Thonge,hidu,Maratha

	//Split ->[]string

	str := "Jay Jay Shree Ram"
	strSli2 := strings.Split(str, " ") // return the slice ,have to pass the string and split wrid
	fmt.Println(strSli2)               //[Jay Jay Shree Ram]

		
}
