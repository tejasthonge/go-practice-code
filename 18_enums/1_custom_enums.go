package main

import "fmt"

//in go enums are not present but we can create by using the const or iota

//in go we can create custom type as follow

type orderStatus string //this is custom type


const (
	Pending  orderStatus= "pending"
	Prepering orderStatus = "prepering"
	OntheWay orderStatus = "ontheway"
	Recevid orderStatus ="recevid"
)

func main(){
	fmt.Println("order Statuse:",Pending)
	fmt.Println("order Statuse:",Recevid)
}

//this is used in project to reduce type o error or maninte the single chage all over the app
