package main

import (
	"fmt"
	"sync"
	"time"
)

func getData(wg *sync.WaitGroup, ch chan []map[string]any) {
	defer wg.Done()
	userData := []map[string]any{
		{
			"id":   1,
			"name": "Tejas",
		},
		{
			"id":   1,
			"name": "Tejas",
		},
		{
			"id":   1,
			"name": "Tejas",
		},
	}
	fmt.Println("getting data ...")

	time.Sleep(3 * time.Second)
	// ch <- userData //blocking
	// we have to run it in anthor go rutine either we get deadlock error
	//as bellow
	go func() {
		ch <- userData
	}() //blocing

}

func main() {
	var wg sync.WaitGroup
	var ch chan []map[string]any = make(chan []map[string]any)
	wg.Add(1)
	go getData(&wg, ch)
	wg.Wait()
	resData := <-ch
	close(ch)
	for _, user := range resData {
		fmt.Println(user)
	}

}
