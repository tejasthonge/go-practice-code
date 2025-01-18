package main

import (
	"fmt"
)

// 1 no parmaetor no return
func demo() {
	fmt.Println("in demo")
}

// 2 parametors
func gun(target string) {
	fmt.Printf("target is %s", target)
}

//3 muiltiple parametors

func playrs(p1, p2, p3 string) {
	con := fmt.Sprintf("playrs are \np1 : %s\np2 : %s\np3 : %s", p1, p2, p3)
	fmt.Println(con)
}

// 4 returns value
func getGun() string {
	return "AK47"
}

//5 multiple retuns values

func getGuns() (string, string, int) {
	return "M416", "AKM", 2
}

// 6 funcion as parametor

func funFuncPara(f func() (string, string, int)) {
	var g1, g2, gCount = f()
	con := fmt.Sprintf("gun count %d ,g1 :%s ,g2 :%s ", gCount, g1, g2)
	fmt.Println(con)
}

// 7 returning the funcion form the funcion callad as closer

func retFun() func() string {
	reFn := func() string {
		return "This return new function"
	}
	return reFn
}

// # this is main entry poind fucion
func main() {
	// 1 no parmaetor no return
	demo()

	// 2 parametors
	gun("pakistan")
	fmt.Println()

	//3 muiltiple parametors
	playrs("Tejas", "Amar", "sayam")

	//4 returns value
	gun := getGun()
	fmt.Printf("Getteing gun : %s", gun)
	fmt.Println()
	//5 multiple retuns values
	var g1, g2, gCount = getGuns()
	con := fmt.Sprintf("gun count %d ,g1 :%s ,g2 :%s ", gCount, g1, g2)
	fmt.Println(con)

	// 6 func as parametor
	funFuncPara(getGuns)

	//7 returning function form function
	ret := retFun()
	resStr := ret()
	fmt.Println(resStr)

}
/*

OP:
in demo
target is pakistan
playrs are 
p1 : Tejas
p2 : Amar
p3 : sayam
Getteing gun : AK47
gun count 2 ,g1 :M416 ,g2 :AKM 
gun count 2 ,g1 :M416 ,g2 :AKM 
This return new function
*/