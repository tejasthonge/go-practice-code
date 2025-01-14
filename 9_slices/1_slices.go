package main

import "fmt"

//slices
//	 - are nothing  but dynamic array
//	- at creating type we does not have to difine the size of slice

func main() {

	//we can create slice as as folllow
	// 1.uninitlized slice

	var sliceArr []int //it  is -->nil means null like dart
	fmt.Println(sliceArr)
	fmt.Println(sliceArr == nil)
	fmt.Println(len(sliceArr))

	//2 intionluzed slice
	//we have the make function

	var sliceArr2 = make([]int, 3) //it hase intial size of threee
	//that mens we not add the 4 th elelnt it is false we can also add the
	//means it is growable array like list in dart

	fmt.Println(sliceArr2)
	fmt.Println(sliceArr2 == nil)

	var slicArr3 = make([]int, 2, 5)
	fmt.Println(cap(slicArr3)) //5
	//here we just defining the capacity
	// it is dynaic array so size will incgess automaticaly
	fmt.Println(len(slicArr3)) //2
	fmt.Println(slicArr3)      //[0 0]
	//it also growable

	//adding the elemet
	//append

	slicArr3 = append(slicArr3, 9)
	/*
		slice = append(slice, elem1, elem2)
		slice = append(slice, anotherSlice...)
	*/
	fmt.Println(slicArr3) //[0 0 9]
	//the element was added in last position
	slicArr3 = append(slicArr3, 5)
	slicArr3 = append(slicArr3, 32)
	slicArr3 = append(slicArr3, 532)

	fmt.Println(slicArr3)      //[0 0 9 5 32 532]
	fmt.Println(cap(slicArr3)) //10

	//the elemt len is 6 but capacity was automaticly incrged and it is now 10
	//	 - it follow the algorithum as it will double the current capacity

	//this is due to of make function

	//3 anthor way to create and slice

	sliceArr4 := []int{}
	fmt.Println(sliceArr4)      //[]
	fmt.Println(cap(sliceArr4)) //0

	// #adding elemt by using the index

	// sliceArr4[3] = 234
	fmt.Println(sliceArr4) //panic: runtime error: index out of range [3] with length 0

	var sliceArr5 = make([]int, 4)
	sliceArr5[3] = 234
	fmt.Println(sliceArr5) //[0 0 0 234]

	// this above is wroking the resion we are initize the initial capacity

	// var sliceArr6 = make([]int, 0, 5) //make(slice[] ,len,cap)
	// sliceArr6[3] = 234 // but this gives error the reson of we provide intial lent as 0 but capcity has ok

	//#copy

	// - if we want copy one slice to anthor we use copy funciton in go
	// var slicArr7[]int // this not work
	var slicArr7 = make([]int, len(slicArr3))
	copy(slicArr7, slicArr3) //copy(destination,source)
	fmt.Println(slicArr7)    //[0 0 9 5 32 532]

	//#slice oprtions

	var numSlice = []int{3, 23, 2, 345, 53}
	fmt.Println(numSlice[0:])  //it will preint all the elemte form 0 th index [3 23 2 345 53]
	fmt.Println(numSlice[:4])  //it print all the elment till 5 expent 4 th position [3 23 2 345]
	fmt.Println(numSlice[3:4]) //it will print contin elemet [345]


}
