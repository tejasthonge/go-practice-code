package main

import "fmt"

//considering our app is use the mutliple paymentgatway for eg.

type Payment struct {
	getWay Ruyzorpay
}

func (p Payment) makePayment(amount float32) {
	p.getWay.pay(amount)
}

type Ruyzorpay struct{}

func (gt Ruyzorpay) pay(amount float32) {
	fmt.Println("Making payment by using Ruzorpay of amout", amount)
}
func main() {
	var amount float32 = 999.00
	var newPayment Payment = Payment{}
	newPayment.makePayment(amount)

}


//supose we have to add anothor payment getway 
//then check next no module