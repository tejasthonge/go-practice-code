package main 

import "fmt"

const (
	zero int = iota
	one int = iota
	two int = iota
	three int = iota
)

func main(){
	fmt.Println(zero)
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
}

/*
0
1
2
3
*/

//iota :
	//ot is auto increment int strinting form zero