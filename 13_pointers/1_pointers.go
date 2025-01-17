package main

import "fmt"

func main() {

	a := 34
	copyA := a       //here he was toring the actule values of a that is 34
	addressOfA := &a //here it storing the address of a

	fmt.Println(copyA)      //34
	fmt.Println(addressOfA) //0xc0000120f0

	//*
	//--if we wanat prinnt the actule value of ginve pointer then we use * as
	fmt.Println(*addressOfA) //34
	fmt.Println(&addressOfA) //0xc00005c030

	addressOfAPtr := &addressOfA // now we storing the address or catual refance
	fmt.Println(addressOfAPtr)   //0xc00005c030
	fmt.Println(&addressOfAPtr)  //0xc000094038
	fmt.Println(*addressOfAPtr)  //0xc0000120f0
	fmt.Println(**addressOfAPtr) //34

	/*
		34
		0xc000090020
		34
		0xc000094028
		0xc000094028
		0xc000094038
		0xc000090020
		34
	*/
	*addressOfA = 80 //here we have the address of a in addressOfA but we changeing thire value as * so that addree storing now thhe 80
	fmt.Println(a)   //80
	// fmt.Println(*a)  //it storing the actule value and so that we can not access valure form alue
}

/*

	& -it is used for the geting the adrress
	* - it is used for geting the value of that adrees

*/
