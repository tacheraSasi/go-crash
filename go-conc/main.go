package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(500 * time.Millisecond) 
	}
}

func printLetters() {
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Printf("Letter: %c\n", ch)
		time.Sleep(700 * time.Millisecond) 
	}
}

func main()  {
	// Start goroutines
	go printNumbers()
	go printLetters()

	// Wait for the goroutines to finish (in a real program, use sync.WaitGroup or channels)
	time.Sleep(4 * time.Second)

	fmt.Println("All tasks completed!")
}