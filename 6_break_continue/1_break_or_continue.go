package main

import "fmt"

func main() {

	//break
	fmt.Println("break statement")

	for i := 0; i <= 10; i++ {

		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("continue statement")

	//continue
	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

/*
op:
break statement
0
1
2
3
4
continue statement
0
2
4
6
8
10

*/
