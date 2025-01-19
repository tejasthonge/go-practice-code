package main

import "fmt"

type Paymenter interface { //this is an interface , the naming conventionl as '...er' at ened of any interce
	pay(amount float32) //this is abostract method of inter case
}

//in other languge impmets key word is present to know which class in implimetint theire interface
//byt in go thire is no like that key word
//it was automaticly know whow interface was implimenting the the interface

type Payment struct {
	// paymetGetway Ruyzorpay
	paymetGetway Paymenter
}

func (p Payment) makePayment(amount float32) {
	p.paymetGetway.pay(amount)
}

type Ruyzorpay struct{}
type Striype struct{}

func (gw Ruyzorpay) pay(amount float32) { //due to pay method sineture as define in paymeter interce is same so Ruyzorpay impements the Paymenter interfase
	fmt.Println("making paymet by using the ruyzorpay ", amount)
}
func (gw Striype) pay(amount float32) { //same as
	fmt.Println("making payment by using Stripe ", amount)
}

func main() {
	var amunt float32 = 999

	//instance of both of gewway
	var ruyzorpayGw Ruyzorpay = Ruyzorpay{}
	var stripeGw Striype = Striype{}

	var newPayment1 Payment = Payment{
		paymetGetway: stripeGw, //here we passing the Striype instance
	}
	var newPayment2 Payment = Payment{
		paymetGetway: ruyzorpayGw, //here we passing the ruzorpay instance
	}
	newPayment1.makePayment(amunt)
	newPayment2.makePayment(amunt)

	//now above withod change the logic we just change what type of is pasing dpending ont that itwill make the payment 
	//by uing the interface 

}

/*
OP:
making payment by using Stripe  999
making paymet by using the ruyzorpay  999


*/