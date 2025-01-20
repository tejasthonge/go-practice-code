package main

import (
	"fmt"
	"sync"
	"time"
)

func getUser(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("geting user , wait 3 sec....")
	time.Sleep(time.Second *3)
}

func main(){
	var wg sync.WaitGroup

	fmt.Println("Start main")
	wg.Add(1)
	go getUser(&wg)
	wg.Wait()
	fmt.Println("End main")
}

/*
Start main
geting user, wait ....
End main

*/