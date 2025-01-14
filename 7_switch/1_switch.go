package main

import (
	"fmt"
	"time" //it is inbuilt labolary
)

func main() {

	//simple swith
	i := 3
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("other number")
	}

	day := time.Now().Weekday()

	switch day {
	case time.Monday:
		fmt.Println(time.Monday)
	case time.Wednesday:
		fmt.Println(time.Wednesday)
	case time.Friday:
		fmt.Println(time.Friday)
	default:
		fmt.Println("day is :", day)
	}

	//multiple condicton switch statement
	switch day {
	case time.Sunday, time.Saturday: //by using the , we can add the multipe conditons on the statement
		fmt.Println("its wekend and day is :", day)
	default:
		fmt.Println("Workdary and day is :", day)
	}

	//type swith

	whoAmi := func(i interface{}) {
		switch i.(type) {
		case string:
			fmt.Println(i, "it is string")
		case int:
			fmt.Println(i, "it is  int")
		case bool:
			fmt.Println(i, "it is boolean")
		case float32:
			fmt.Println(i, "it is float32")
		case float64:
			fmt.Println(i, "it is float64")
		default:
			fmt.Println(i, "it is other type")
		}
	}

	whoAmi(2.2)
	whoAmi(true)
	whoAmi(123i)

}

/*
OP: 

three
Monday
Workdary and day is : Monday
2.2 it is float64
true it is boolean
(0+123i) it is other type
*/
