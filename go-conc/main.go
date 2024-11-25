package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(500 * time.Millisecond) // Simulate some work
	}
}

func printLetters() {
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Printf("Letter: %c\n", ch)
		time.Sleep(700 * time.Millisecond) // Simulate some work
	}
}

func main()  {
	
}