

package main
import "fmt"

func main(){

	const email = "amarthonge676@gmail.com";
	const name string = "tejas thonge"

	// email = 344;  here we gtting the error the reson of we cant assing new value constants after we decale ons
	fmt.Println(email);
	fmt.Println(name)


	//we can also use pair of const as 
	const (
		a = 10;
		b string = "stings const"
	)

	fmt.Println(a)
	fmt.Println(b)
}
