package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")

	// get command line arguments
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}

	os.Exit(0)
}
