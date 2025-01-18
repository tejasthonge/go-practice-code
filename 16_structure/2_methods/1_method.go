package main

import (
	"fmt"
)

type Company struct {
	name     string
	empCount int16
}

func (c Company) IncreegeEmpCout() { //here no afet on it
	c.empCount++
}


func (c *Company) IncmetEmpCout() {  //we geting address of instance of the Company
	c.empCount++//her stuct atomatic convert that is derefaning  them as value
}


func main() {

	myComp := Company{
		name:     "Sitaram",
		empCount: 0,
	}
	fmt.Println(myComp)
	myComp.IncreegeEmpCout()
	fmt.Println(myComp) //myComp not incged the count the resoin is above at time of creation of method
	//that we geting the value of our comp not the address so it change on tis own  to change that we have pass the address of
	myComp.IncmetEmpCout()   //but here we incgess actulle cout by pasing thire 
	fmt.Println(myComp)
	//so any modiction on the stuct by the method wecreate method like secton

}
