package main

import "fmt"

//Embedded struct : it is nothong but struct inside other struct

type Employee struct{
	id int
	name string
}

type Company struct{
	name string
	employee Employee
}



func main(){

	var emp1 Employee = Employee{
		id :1,
		name:"Tejas",
	}
	
	fmt.Println(fmt.Sprintf("emplye : %v" ,emp1)) //emplye : %v{1 Tejas}

	var companyInfo Company =Company{ 
		name:"Sitaram",
		employee: emp1,
	}


	fmt.Println(fmt.Sprintf("Compny info : %v",companyInfo ))//Compny info : {Sitaram {1 Tejas}}
	
}