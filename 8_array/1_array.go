package main

import "fmt"

func main() {

	// array :
	//	-- in go array are similar type elemet are contin
	// -- it is numberd sequince of fixed length
	// -- array is hase fixed length we have to defile at time of creating the array

	var arr1 [5]int

	fmt.Println(arr1)

	//len function
	// it is and inbuild funcitoon in go to find the lenth
	fmt.Println(len(arr1))

	/*
	   op:
	   [0 0 0 0 0]
	   5
	*/

	//form the above out put we can say that the defoutl value of each elemnt in array is 0

	// ading and returving elemet form the array

	arr1[0] = 2
	fmt.Println(arr1[0]) //op: 2

	fmt.Println(arr1) //op : [2 0 0 0 0]

	//array of string

	var arr2 [3]string
	fmt.Println(arr2) //op:[  ]

	var arr3 [4]float32
	fmt.Println(arr3) //OP :[0 0 0 0]

	var arr4 [5]bool
	fmt.Println(arr4) //OP :[false false false false false]
	fmt.Println("----------------")

	//itarating on an array
	for i := range arr1 {
		fmt.Println(arr1[i])
	}

	//crating array as well defining the values

	arr5 := [5]int{23, 323, 23, 535, 234}
	fmt.Println(arr5)

	/*
	Limitations
Fixed Size: Once declared, the size of an array cannot be changed.
Use Slices for Flexibility: Slices are preferred in Go for dynamic collections.
	*/
}
