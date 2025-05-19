package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Use: ./ErrorCorrection -c/-d text.txt")

		return
	}

	if args[1] == "-c" {
		// Criptografy Mode
		fmt.Println("Criptografy Mode")
	} else if args[1] == "-d" {
		// Decriptografy Mode
		fmt.Println("Decriptografy Mode")
	} else {
		fmt.Println("Mode Unwknow")
		return
	}

	file, err := os.Open(args[2])
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Content:")
	fmt.Println(string(content))
}
