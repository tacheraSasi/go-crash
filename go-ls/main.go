package main

import (
	"fmt"
	"os"
)

//This is a simple implemenatation of the ls command in linux
//Built with the golang programming language by Tachera W




func main()  {
	dir := "."
	if len(os.Args) > 1 {
		fmt.Println(os.Args[0])//for curiosity only 
		dir = os.Args[1]
	}

	
}