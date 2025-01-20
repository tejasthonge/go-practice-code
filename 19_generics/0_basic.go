// after 1.18 version
package main

import "fmt"

func printNumSlice(items []int) {
	for _, val := range items {
		fmt.Println(val)
	}
}

func printSlice[T any](items []T) { //here it ressicre any type of slice ,and now T is templete type  ,,we can use the interface{} replace of any 
	for _, val := range items {
		fmt.Println(val)
	}
}

func main() {
	var nums []int = []int{234, 342, 4232, 24}
	var forts []string = []string{"Shivneri", "Rajgad", "Raigad"}

	printNumSlice(nums) //but this function not genaric means it only print the int types
	// so that time come the new featur callad as generic
	// printNumSlice(forts) // it  well getting the error the resion of  it can resive only integer type slice only

	// but anthor function we created by using the gererinc type

	printSlice(nums)  // here we pass the int type
	printSlice(forts) //and here  we pass the string type but also these work
	//like that we use the generic
	//genrics on map

	var mapEg map[string]string = map[string]string{ //this type we provide the generic that is key is string and value is also string
		"name": "tejas",
	}

	fmt.Println(mapEg)

}
