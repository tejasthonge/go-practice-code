package main

import "fmt"

//structure:
/*
	- in go thire no class and object concept so we have the fill like class or object or
	- to overcome the class functionalty in go
	- this structuture comes
	- the key word used for it is struct
*/

type Persion struct {
	name    string 
	age     int
	address string
}

func main() {

	var persion1 Persion = Persion{
		name:    "Tejas",
		age:     23,
		address: "pune",
	}

	persion2:= Persion{
		name: "Amar",
		age:23,
		address: "Barshi",
	}

	fmt.Println(persion1) //{Tejas 23 pune}
	fmt.Println(persion2) //{Amar 23 Barshi}

	
	//acassing single elemet

	name := persion1.name
	fmt.Println(name)
	fmt.Println(persion1.age)
	fmt.Println(persion1.address)


	persion1.address = "Bangalore"
	fmt.Println(persion1.address)

	fmt.Println(persion1)


}
