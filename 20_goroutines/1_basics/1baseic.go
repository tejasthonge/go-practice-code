package main

import (
	"fmt"
)

//go :

// go langauge is multithreaded language that have the
// by using theis we can muiltiple fuctions or possess at a time conquirently
// go create the lightwight threads

func haveyTask() {
	for i := 0; i < 10000000000000000; i++ {
		fmt.Println("doing task", i)
	}
}

func main() {
	fmt.Println("Strati main")
	go haveyTask()
	fmt.Println("end main")

	/*
		Strati main
		end main
		doing task 0
	*/
	//form the above out put the main thread run securnsy but 
	//due go key word it run in anthor thread 
	// all out put not printed the resion of till they pinting main thread is ended
}
