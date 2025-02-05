package main

import "fmt"

func main() {

	//int -- to store the numic values
	// it hase three type sub types or ..
	//1 int64
	//2 int32
	//3 int8

	var accountBalance1 int = 2300000000000000000   //must contain 18 digit it is by defoult int64 bit
	var accountBalance2 int64 = 2300000000000000000 //it is same as defoult int
	var phoneNumber int32 = 755866798               //it can contain only 9 digit
	var pin2 rune =421619123// same as the int32
	var pin int16 = 6767

	fmt.Println(accountBalance1)
	fmt.Println(accountBalance2)
	fmt.Println(phoneNumber)
	fmt.Println(pin)
	fmt.Println(pin2)
	
	// we can store all this above values in int
	//go intenly mangae the storge automaticly
	var nagVal int = -31
	fmt.Println(nagVal)
}
