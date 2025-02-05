package main

import (
	"fmt"
	"strings"
)

//strings packge
//ackage strings implements simple functions to manipulate UTF-8 encoded strings.

func main() {

	str1 := "Tejas"

	fmt.Println(strings.ToLower(str1))
	fmt.Println(strings.ToUpper(str1))

	fmt.Println(strings.Contains(str1, "ja"))
	fmt.Println(strings.Contains(str1, "KaA"))

	str2 := "Tejas8055"
	fmt.Println(strings.Compare(str1, str2)) //it return the int value as

	fmt.Println(strings.ContainsRune(str2, 0))

	//trim

	strTr := "    fdsfdfdsf  "
	fmt.Println(strings.Trim(strTr, " ")) //here it was triming the spaces but we rpassing them as latter f then
	fmt.Println(strings.Trim(strTr,"f"))

}
