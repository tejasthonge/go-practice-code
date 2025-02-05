package main

import (
	"fmt"
	"strconv"
)

func main() {

	fString := strconv.FormatBool(true)
	fmt.Println(fString)

	strFloat:=strconv.FormatFloat(32.423,64,32,33)
	fmt.Println(strFloat)

}
