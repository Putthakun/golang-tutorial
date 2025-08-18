package main

import "fmt"

type Customer struct {
	ID string
	Name string
}

func (c Customer) checkCustomer() string {
	if c.ID == "" { return "No member" }
	return c.ID + " " + c.Name
}

func main() {

	// create struct
	type product struct {
		name string
		price int
		stock int
	}

	// how to use struct
	var p1 product

	p1.name = "bag"
	p1.price = 299
	p1.stock = 20

	fmt.Println(p1)

	//acess 1 item
	fmt.Println(p1.name)

	// ------------------------ medthod and struct -------------------

	customer1 := Customer{
		ID : "1234",
		Name : "Putthakun",
	}

	//func(receiver) Method

	// Call method
	fuoundCustomer := customer1.checkCustomer()
	fmt.Println(fuoundCustomer)
}

