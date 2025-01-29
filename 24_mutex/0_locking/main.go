package main

import (
	"fmt"
	"sync"
)

func proseesBigData() {
	fmt.Println("Prossing big data..")
}

func main() {
	m := new(sync.Mutex)
	fmt.Println("Starting main")
	go proseesBigData()
	m.Lock()
	fmt.Println("End Main")

}
