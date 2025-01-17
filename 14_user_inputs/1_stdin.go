package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Inter the input")

	reader := bufio.NewReader(os.Stdin) //here we create the input reader it hase methods like
	// reader.ReadLine();
	// reader.ReadString('\n')
	// reader.ReadRune();
	//or more methods are present in it

	input, err := reader.ReadString('\n')
	if err == nil {
		fmt.Println(input)
	}
	fmt.Println("enter int value")

	numStr, err := reader.ReadString('\n')
	numInt, err := strconv.ParseInt(strings.TrimSpace(numStr), 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		conInt := int(numInt)
		concStr := fmt.Sprintf("you are ented int %d", conInt)
		fmt.Println(concStr)
	}

}
