package main

import "fmt"

func main(){
	var names = []string{"Irene","Rachel","Agatha","Stella"}

	//The normal old fashion for loop
	for i := 0; i < len(names); i++{
		fmt.Println("This girl broke my heart",names[i])
	}
}