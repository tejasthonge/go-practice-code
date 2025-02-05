// Finding and Replacing Substrings

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is string athat we have to replase and we will to by using the method and it esy to do in go and"
	fmt.Println(str)
	strRep := strings.Replace(str, "and", "or", -1)
	fmt.Println(strRep)

	fmt.Println(strings.Index(str, "and"))
	/*
		This is string athat we have to replase and we will to by using the method and it esy to do in go
		This is string athat we have to replase or we will to by using the method or it esy to do in go
		40
	*/

	

}
