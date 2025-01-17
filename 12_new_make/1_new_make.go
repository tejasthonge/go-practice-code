package main

import "fmt"

func main() {

	var arr = [5]int{}
	fmt.Println(arr) //[0 0 0 0 0] defolat values are 0

	// createing arry by using the make and new and seing the conclusion
	var arr2 = make([]int, 4)
	fmt.Println(arr2) //[0 0 0 0]

	var arr3 = new([3]int)
	fmt.Println(arr3)  //&[0 0 0]
	fmt.Println(*arr3) //[0 0 0]
	fmt.Println(&arr3) //0xc00005c038

	//make :-return acture value
	//new :- it was return  the refarance or pointer of any values

	fmt.Println("-------------")
	//make
	//	-- it allocate the memory alos intilize
	//	-- you will get the memory address also
	//  -- non sezo stored

	// new
	// 	-- allocated the memory so no init
	// -- so it returning the adrees only
	//  -- zero stored

	//for ex:

	var a = new(int)
	fmt.Println(a)  //0xc000104ed0
	fmt.Println(*a) //0

}
