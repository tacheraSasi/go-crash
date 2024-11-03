package ranges

import "fmt"

func main(){
	var names []string = []string{"Irene","Rachel","Agatha","Stella"}

	//The normal old fashion for loop
	for i := 0; i < len(names); i++{
		fmt.Println("This girl broke my heart",names[i])
	}

	fmt.Println("\nUsing the range keyword instead")
	fmt.Print("-------------------------------\n")

	for i,name := range names{
		fmt.Println(i+1,"Thiis girl also broke my heart",name)
	}
	
}