package main

import "fmt"

type Point struct{
	x int32
	y int32
}
// The function to change the value of x in a pointer
//From the default using an pointer 

func main(){
	p1 := &Point{y:3}//only giving the value to one item in the struct
	//Leaving the other one to take the default value of its type

	fmt.Println(p1)
	change(p1)
	fmt.Println(p1)

	
}