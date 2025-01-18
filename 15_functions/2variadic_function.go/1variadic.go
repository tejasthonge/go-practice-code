package main

import "fmt"

func demo(nums ...int){ //nums is now slice
	for _,n := range nums{
		fmt.Println(n)
	}

}

func main(){
	fmt.Println("fd",212,1,"dfd")//here we pasing the multiple favluse
	//So Println fuction callad as variadic fucntion

	demo(23,212,42,1)
	demo(5)
}
/*
OP:
fd 212 1 dfd
23
212
42
1
5 
*/