package mapModule

import "fmt"

func Map() {
	// create map using make()
	m := make(map[string]int)  // key = string, value = int
	m["apple"] = 5
	m["banana"] = 6

	fmt.Println(m)

	// map interal
	m2 := map[string]int {
		"adam" : 20,
		"json" : 35,
	}

	fmt.Println(m2)

	// access Map value
	fmt.Println(m["apple"])
	fmt.Println(m["adam"])

	// check key
	val, ok := m["apple"]  // val = value, of = true or false
	if ok {
		fmt.Println("found: ", val)
	}else {
		fmt.Println("not found")
	}

	// delete map 
	delete(m, "banana")
	fmt.Println(m)

	// Loop in Map
	for k, v := range m2 {  // unordered
		fmt.Println(k, v)  // k = key, v = value
	}

	// lenght Map
	fmt.Println(len(m2))

	// nil Map
	var mNil map[string]int
	fmt.Println(mNil == nil)

	// make
	mNil = make(map[string]int)
	mNil["doge coin"] = 6
	fmt.Println(mNil)


}