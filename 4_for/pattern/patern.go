package main

import "fmt"

func main() {
	/*

	*
	**
	***
	****
	*****
	****
	***
	**
	*
	 */

	var k int = 1

	for i := 1; i < 10; i++ {
		for j := 1; j <= k; j++ {

			fmt.Print("*")
		}
		if k < 5 && i < 5 {
			k++
		} else {
			k--
		}

		fmt.Println()
	}

	//------------------

	/*
	   	   *
	      ***
	     *****
	    *******
	   *********
	    *******
	     *****
	      ***
	   	   *

	*/

	// var mid = 5
	// var row = 10

	// for i := 1; i < row; i++ {

	// 	for j := 1; j < row; j++ {
	// 		if i <= mid {
	// 			if j <= i && j >= mid-i {
	// 				fmt.Print("*")
	// 			} else {
	// 				fmt.Print(" ")
	// 			}
	// 		} else if i >= mid {

	// 		} else {
	// 			fmt.Print(" ")

	// 		}
	// 	}
	// 	fmt.Println()
	// }

}
