package template

import "fmt"

type Customer struct {
	createInGateway func()
}

func (r *Customer) Create(){
	r.createInDatabase()
	r.createInGateway()
}

func (r *Customer) createInDatabase(){
	fmt.Println("Customer created in database")
}

func createInStripe(){
	fmt.Println("Customer created in stripe")
}

func createInPaypal(){
	fmt.Println("Customer created in paypal")
}

func StripeCustomer() *Customer{
	return &Customer{createInStripe}
}

func PaypalCustomer() *Customer{
	return &Customer{createInPaypal}
}
