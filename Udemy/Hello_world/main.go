package main

import (
	"fmt"
	"os"
)

func main() {
	// var args []string
	args := os.Args
	if len(args)<2 {
		fmt.Printf("Usage: ./hello-world <argument>\n")
		os.Exit(1)
	}
	fmt.Printf("Hello!\nos.arg: %v\n1st Argument: %v\n", args, args[1:])
}
