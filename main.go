package main

import (
	"dirscan/scanner"
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args

	argsAmount := len(argsWithProg)
	if argsAmount < 2 {
		fmt.Println("Please provide a file to search for.")
		return
	} else if argsAmount > 2 {
		fmt.Println("It will only look for the first file name provided.")
		return
	}
	fileToFind := argsWithProg[1]
	scanner.DirScan(fileToFind)
}
