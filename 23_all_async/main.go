package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// at time of marshing in to json we must have to use the tags as bellow
type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"role"`
}

func getUserInfo(id int) (User, error) {
	fmt.Println("Getting user info ..")
	time.Sleep(time.Second * 2)
	return User{
		Id:   id,
		Name: "Tejas Thonge ",
		Age:  23,
		Role: "Bakend developer",
	}, nil
}

func getUserIdByRef(wg *sync.WaitGroup, id *int, err *error) (int, error) {
	defer wg.Done()
	fmt.Println("getting user id ..")
	time.Sleep(time.Second * 3)
	*id = 12
	*err = nil
	return 12, nil
}
func getUserIdByChan(wg *sync.WaitGroup, idChan chan int) (int, error) {
	defer wg.Done()
	fmt.Println("getting user id ..")
	time.Sleep(time.Second * 3)
	idChan <- 99
	return 12, nil
}

func main() {
	fmt.Println("start main")
	var wg sync.WaitGroup
	var id int
	var err error
	//without chan
	wg.Add(1)
	go getUserIdByRef(&wg, &id, &err) //here can not get the the return value so we pass the pointeres of new wrible
	wg.Wait()
	//by using chan
	fmt.Println("user id", fmt.Sprintf("id: %d", id))
	idChan := make(chan int)
	wg.Add(1)
	go getUserIdByChan(&wg, idChan)
	wg.Wait()
	id = <-idChan
	close(idChan)

	fmt.Println("user id", fmt.Sprintf("id: %d", id))

	if err != nil {
		fmt.Println("Getting an error at geteeing id")
		return
	}

	userInfo, err := getUserInfo(id)
	if err != nil {
		fmt.Println("getting error at time feathing user info")
	}

	userJsonBytes, err := json.Marshal(userInfo) //it return the json or error
	if err != nil {
		fmt.Println("getting the error to marshing struct to json")
	}
	fmt.Println(string(userJsonBytes))
	fmt.Println("start end")

}
