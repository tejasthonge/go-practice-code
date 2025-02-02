package main

import (
	"fmt"
)

/*
Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
*/
func main() {

	fmt.Println("This is fmt println func")

	name := "Tejas "
	comp := "qbitrom"
	role := "sopftwer devloper"
	age := 23
	//  strings conatination
	conStr := fmt.Sprintf("name : %s ,role : %s company %s age %d", name, role, comp, age) // it return the string

	fmt.Println(conStr)
	// hase other method like
	// var w io.Writer
	// fmt.Fprint(w, "thonge") //t
	byt := []byte("this the bytes")

	apByte := fmt.Append(byt)
	fmt.Println(apByte) //[116 104 105 115 32 116 104 101 32 98 121 116 101 115]

	err:=fmt.Errorf("This is error maessage")  //it erturn the Error
	fmt.Println(err.Error())
	


	
}
