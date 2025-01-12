

package main

import "fmt"

//in go only have the for loop 

//thire is no while dowhile like other porgaming languge

func main(){
	var i int8 = 0;

	//while
	//the following code woring as while 
	for i<10{
		fmt.Println("value of i :" ,i);
		i++
	}

	


	//classic for loop
	for j :=0;  j<=10; j++{
		fmt.Println("value of j:",j)

	}
	// infinite loop
	
	for {
		fmt.Println("runing inpinte time")
	}
	
}