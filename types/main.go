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
//Anything from the actual point
func (p Point) getPointsInAnArray() []int32 {
	arr := []int32{p.x,p.y}
	return arr

}

//Struct method that change the values in the actual point
//It needs to recieve a pointer to the struct so that i can 
//Modify it instead of making a copy
func (p *Point )changeVal(x,y int32){
	p.x = x
	p.y = y

}

func main(){
	p1 := &Point{y:3}//only giving the value to one item in the struct
	//Leaving the other one to take the default value of its type

	fmt.Println(p1)
	changeX(p1,100)
	fmt.Println(p1)

	//printing out the array of the points
	//From the Point struct method
	fmt.Println(p1.getPointsInAnArray())

	//Changing then calues of the point using the struct method
	p1.changeVal(34,65)
	//Then Printing out their values once again
	fmt.Println(p1)


	
}