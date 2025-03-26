package main

import (
	"fmt"
	"os"
)

func main() {
	// Print greeting message with program name
	if len(os.Args) > 1 {
		fmt.Printf("Hello, %s!\n", os.Args[1])
	} else {
		fmt.Println("Hello, World!")
	}
}