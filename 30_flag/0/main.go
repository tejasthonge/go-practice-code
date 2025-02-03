package main

/*

The flag package enables command-line flag parsing in Go programs, allowing developers to easily define and handle command-line arguments.
*/
import (
	"flag"
	"fmt"
)

func main() {

	host := flag.String("host", "localhost", "Put your database host") //here we have to pass the three string first is the falg name ,second is falg values , and the thired is maessage
	port := flag.String("port", ":999", "Put your port")               //it return the address of the variable so we haveto get as the pointer
	flag.Parse()                                                       //if we not parse them it well geting the defoalt values
	conStr := fmt.Sprintf("host : %s ,prort : %s", *host, *port)
	fmt.Println(conStr) //host : localhost ,prort : :8055 geting this when we not setting it mannully else
	// what we passad at time of running the code like
	// go run 30_flag/0/main.go -host="lhost" -port="8055"
	// host : lhost ,prort : 8055

	isDev := flag.Bool("isDev", true, "Put your mode")
	fmt.Println(*isDev)
	//go run 30_flag/0/main.go -host="lhost" -port="8055" isDev=false
	//true

	/*
		flag.String(): String value
		flag.Int(): Integer value
		flag.Bool(): Boolean flag
		flag.Float64(): Floating-point number
		flag.Duration(): Time duration
	*/

	

}
