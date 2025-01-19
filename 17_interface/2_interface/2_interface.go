package main

import "fmt"


type Payment struct {
	// paymetGetway Ruyzorpay 
	paymetGetway Striype 
	
}
func (p Payment)makePayment(amount float32){  //this method of Payment, for mking the paymetn
	p.paymetGetway.pay(amount)
}

//this are two getway 
type Ruyzorpay struct{}  
type Striype struct{}
//and eacha have itw own pay method e

func (gw Ruyzorpay)pay(amount float32){
	fmt.Println("making paymet by using the ruyzorpay " ,amount)
}
func (gw Striype)pay(amount float32){
	fmt.Println("making payment by using Stripe ",amount)
}

func main(){
	var amunt float32 = 999;

	//instance of both of gewway
	// var ruyzorpayGw Ruyzorpay =Ruyzorpay{}
	var stripeGw Striype =Striype{}

	var newPayment  Payment =Payment{ 
		paymetGetway: stripeGw,  //here we passing the ruzorpay instance 
	}
	newPayment.makePayment(amunt)  

	//but this above case we have change the payment getw by hardcoding hard codeing 
	//as seening the above code 

	//till intercace is not comming 
}