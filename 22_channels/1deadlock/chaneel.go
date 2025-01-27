//chaneels :

/*
	channels are used share the data form one go routene to anthor go rutene by using this we
	can comminicate bettwern two go rutines
	-
*/

package main

import "fmt"

func main() {

	// by these way we can create the channel by using the make fucntion by passing chan keyword and which
	//type of data we are used commincate bettwen the go rutines

	messsageChan := make(chan string)

	//sending data in channel

	messsageChan <- "DataMessage" //her we inserting the data in cheellal by using the <- and direction of arrow towods the chen 

	//ressing or getting the data form channer
	resivedData := <-messsageChan   //here we ressiving the data form chan and the direction of thes <- data form chan to varible 
	fmt.Println(resivedData)
	
	/*
	op:
			fatal error: all goroutines are asleep - deadlock!

			goroutine 1 [chan send]:
			main.main()
				/Users/tejasthonge/coding/github/go-practice-code/22channels/1basics/chaneel.go:22 +0x36
			exit status 2
	*/

	//this error is due to of

	/*
		in our case 
		wehen we pass the data that time the due to chennal it will bocking till it will ressive the data
		so it is blocking then 
		next line of code not exucute till it will relese or unblock 
		but to release or unblock it will have to resive the data form chan 

		so this deadlock is happaning
	*/
	/*
	

	deadlock:
		her two process anre runing symaltenlly and fist ruing and wait for the getting ressors and end the task
		and anthor prosses waiting for reles fist prossen and run second mens owin prosses 
		but here 
		both are wating and not complitin or relsesing so that thes err is got

		mens pro1 waitnig for data1 
		and this data or resorse is coming form 
		pro2 but but 
		pro2 two wating for exucite thire own prosses but pro1  is renuing so pro2 witing for it 
		means buth are run infinte time loop so that is knon as deadlock 

	*/
}
