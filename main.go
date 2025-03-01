package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	fileToFind := argsWithProg[1]
	fmt.Println("Searching for file:", fileToFind)
	fmt.Println("Hello, World!")
}
