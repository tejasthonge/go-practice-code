package main

import "fmt"

func main() {

	//creating the map without make function

	mapEg :=map[string]string{
		"rajy":"Akhand Bharat",
		"king":"Shree Ram",

	}
	fmt.Println(mapEg)  //map[king:Shree Ram rajy:Akhand Bharat]

	//in go we can return fultiple value 
	//so 
	// var _, ok = mapEg["pap"];
	//or 
	_,ok := mapEg["pap"]  //it return the value ,isPresent  keyo
	//we using the _ so we not used it then also not got error but replase of that any varible we use it geting error
	if ok{
		fmt.Println("Present")
	}else{
		fmt.Println("Not Present") 
	}
	


	fmt.Println("--------")
	//createing map by using make function
	
	var m map[string]any = make(map[string]any)
	m["name"] = "Tejas"

	fmt.Println(m) //map[name:Tejas]
	fmt.Println(m["name"])
	m["role"] = "Softweare Developer"
	m["age"] = 23
	fmt.Println("----")
	for k := range m {
		fmt.Println(k) //keys
	}
	fmt.Println("------")
	for k := range m {
		fmt.Println(m[k]) //valuse
	}

	//accasing not exiting key
	fmt.Println(m["addr"]) //<nil>

	for key, val := range m {
		r := fmt.Sprintf("%s : %v", key, val) //stirng interpolation
		fmt.Println(r)
	}

	//deleting or removing elemet form map

	delete(m,"age");
	for  k ,val:= range(m){
		stringCon:= fmt.Sprintf("%s : %v , ",k,val);
		fmt.Printf("%s", stringCon) //name : Tejas , role : Softweare Developer ,  
	}

	
}
