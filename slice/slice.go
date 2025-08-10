package slice 

import "fmt"

func Slice() {
	
	// crate slice
	names :=[] string {"สมชาย", "แก้วตา"}
	fmt.Println("Before add data: ", names)
	fmt.Println("Before len of data: ", len(names))
	fmt.Println("Before cap of  data: ", cap(names))

	// add value in slice
	names = append(names,"แดงดำ")
	fmt.Println("After add data: ", names)
	fmt.Println("After len of data: ", len(names))
	fmt.Println("After cap of data: ", cap(names))

	// Access data 
	fmt.Println("Acess index 0: ", names[0])

	// split
	subNames := names[0:2]
	fmt.Println("Split slice names index 0 --> index 1: ", subNames)

	//reference type
	price := []int {10, 20, 30}
	copyPrice := price[:]
	copyPrice[0] = 60

	fmt.Println("Price index 0:", price)
	fmt.Println("Copy price index 0:", copyPrice)

	// make()
	s := make([]int, 3, 5) // type, length, capacity
	fmt.Println(s)         // [0 0 0]
	fmt.Println(len(s))    // 3
	fmt.Println(cap(s))    // 5
}
 