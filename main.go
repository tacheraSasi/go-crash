package main

import "fmt"

var print = fmt.Println

func main(){
	var name string = "tachera"
	output := "My name is "+name
	// fmt.Println("my name is",name)
	fmt.Println(output)

	//arrays
	arr := []int{2,3,5,66,8,34,5,6,1,2,23 }
	sum := 0
	for i := 0; i < len(arr);i++{
		sum += arr[i]

	}
	print(&sum)
	print(*&sum)
}
