package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Get the current directory or take from args
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	// Open the directory
	directory, err := os.Open(dir)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer directory.Close()

	// Read the directory contents
	files, err := directory.Readdir(-1) // -1 means read all entries
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Loop through the files and display information
	fmt.Println("Listing files in directory:", dir)
	for _, file := range files {
		// File name
		fmt.Printf("%-25s", file.Name())

		// File size
		fmt.Printf("%-10d", file.Size())

		// File permissions
		fmt.Printf("%-10s", file.Mode().String())

		// Last modified time
		modTime := file.ModTime().Format(time.RFC1123)
		fmt.Printf("%-20s\n", modTime)
	}
}
