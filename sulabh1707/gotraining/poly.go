package main

import (
	"fmt"
)

type Customer interface {
	register() string
	login() string
}

type Buyer struct {
	name string
	id   int
}

type Seller struct {
	name string
}

func (b Buyer) register() string {
	return fmt.Println("You are registered as Buyer")
}

func (b Buyer) login() string {
	return fmt.Println("You are login as Buyer")
}

func (s Seller) register() string {
	return fmt.Println("You are registered as Seller")
}

func (s Seller) login() string {
	return fmt.Println("You are login as Seller")
}

func ride(vehicle Vehicle) {
	fmt.Println(fmt.Sprintf("Literal Vehicle: Value is %v. Type is %T.", vehicle, vehicle))
	fmt.Println(vehicle.description())
	fmt.Println(fmt.Sprintf("Did you notice that the vehicle rides on %v?", vehicle.medium()))
}

func measure(c Customer) {
	fmt.Println(c)
	fmt.Println(c.login())
	fmt.Println(c.register())
}
func main() {
	b := Buyer{name: "sulabh", id: 4}
	s := circle{name: "katiyar"}
	//The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
	measure(b)
	measure(s)
}
