package main

import (
	"fmt"
	"sync"
)

func syncProsses(wg *sync.WaitGroup, msgChan chan string) {
	defer wg.Done()
	var resData string = <-msgChan
	fmt.Println(resData)
}

func main() {
	var wg sync.WaitGroup
	msgChan := make(chan string)
	wg.Add(1)
	go syncProsses(&wg, msgChan)
	msgChan <- "This is data"
	close(msgChan) // we will cosing after that all the data is resived by the func syncProsses
	wg.Wait()

}

//means here we fist create the chan and call th go rutine function
// afther that we pass the data to chan and this will resvide to anthor go rutine that is prossese is running on it after resing the data
// it wll automatcally release exucute the next line
