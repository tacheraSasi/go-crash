package main

import (
"fmt"
)

var print = fmt.Println

func main(){
	// array with the short notation
	arr := [5]int{1,2,3}
	print(arr)
	var array [5]int = [5]int{1,2,3,4,5}// a normal array
	var slice []int = array[:]

	slice = append(slice,6)
	//appending an element to a slice
	slice2 := append(slice,70)

	//printing the copied slice
	print(slice)
	print(slice2)

}
	
