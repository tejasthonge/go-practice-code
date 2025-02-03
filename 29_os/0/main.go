package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	fmt.Println(args) //[/var/folders/t_/vvf4vg6n29d_z5btz5xy7g000000gn/T/go-build1144462560/b001/exe/main]

	//geting Hostname

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Getting the error to get the host name ", err)
	} else {
		fmt.Println(hostname) //boss.local
	}



}
