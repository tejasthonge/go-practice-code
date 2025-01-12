

package main
import "fmt"

func main(){

	const   i int8 = 10;

	//if statements
	if i == 0{
		fmt.Println(i,"is equl to 0")
	}
	if i == 10 {
		fmt.Println(i,"is equl to 10")
	}


	//if-else stamets
	var isLogin bool = true;

	if isLogin{
		fmt.Println("user is loged")
	}else{ // we can not write else to the next line
		fmt.Println("User not loged")
	}

	//if ,else-if 

	if age:=10; i <18{ //here we directly define as we assing the value to the age 
		fmt.Println(age,"age is no consider as adult")
	}else if age >= 18{
		fmt.Println(age,"is adult")
	}
	

}