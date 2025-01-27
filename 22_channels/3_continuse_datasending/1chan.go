package main

import (
	"fmt"
	"sync"
	"time"
)

func prossesData(wg *sync.WaitGroup, ch chan int) {
	defer func() {
		wg.Done()
	}()
	for data := range ch {
		fmt.Println(fmt.Sprintf("Prosses data %d", data))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("Start main")
	var wg sync.WaitGroup
	var ch chan int = make(chan int)
	wg.Add(1)
	go prossesData(&wg, ch)
	for i := 1; i < 10; i++ {
		ch <- i
	}
	close(ch) // we have close after all the data is passing 
	wg.Wait()
	fmt.Println("end main")

}
