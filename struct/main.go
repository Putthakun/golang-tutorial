package main

import "fmt"

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

	fmt.Println(p1.name)
}
