package main

import "fmt"

type Point struct{
	x int32
	y int32
}
// The function to change the value of x in a pointer
//From the default using an pointer 
func changeX(point *Point, xVal int32){
	point.x = xVal//Changing the value of point.x through the actual pointer
}

//struct method that does not change 
//Anythin from the actual point
func (p Point){

}

func main(){
	p1 := &Point{y:3}//only giving the value to one item in the struct
	//Leaving the other one to take the default value of its type

	fmt.Println(p1)
	changeX(p1,100)
	fmt.Println(p1)

	
}