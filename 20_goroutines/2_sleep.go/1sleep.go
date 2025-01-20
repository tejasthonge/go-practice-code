package main

import (
	"fmt"
	"time"
)


func haveyTask(){
	for i:= 0;i<10000;i++{
		fmt.Println("doing task ",i)
	}
}
func main(){

	fmt.Println("start main")
	go haveyTask()
	fmt.Println("end main")
	time.Sleep(time.Second*3)

}	
/*
Strati main
end main
0
1
... all values
*/

// due to sleping the it well print the all vlues till main thread running 
//