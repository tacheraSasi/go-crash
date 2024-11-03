package main

import "fmt"

func main(){
	// The classic way of defining maps
	var map1 map[string]int = map[string]int{
		"apple":2,
		"mangoes":3,
		"oranges":6,
	}

	fmt.Println(map1)
	delete(map1,"apple")
	
	_,ok := map1["apple"]
	if !ok {
		fmt.Println("Apples were deleted sucessfully")
	}
	fmt.Println(map1)
	
}