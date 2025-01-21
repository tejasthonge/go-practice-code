package main

import (
	"fmt"
	"sync"
)

func heveyTask(wg *sync.WaitGroup) {
	defer wg.Done() // due to defer it will after commplite all the prosses 
	for i := 0; i < 100; i++ {
		fmt.Println("doing task :", i)
	}
}

//
func main() {
	var wg sync.WaitGroup
	fmt.Println("Start main")
	wg.Add(1) //here  this will add the 1 tastk or fuction to the shedular then sheduler 
	go heveyTask(&wg)
	wg.Wait() // it will waiting till the hevesTask call the Done method 
	//after calling the done method it will unbock the main thread and exicute next prosees
	fmt.Println("end main")
	
}
