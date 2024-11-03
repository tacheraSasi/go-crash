package main

import "fmt"

// a function that returns a function
func returnFunc(x string) func(){
	return func ()  {
		fmt.Println(x)
	}
}

func main(){
	inLineFunc := func(){
		fmt.Println("This is an in line func similar to arrow function in javascript")
	}
	inLineFunc()

	// Calling the function that returns a function
	returnFunc("This function returns a function")() //immediately revoking it 

	
}