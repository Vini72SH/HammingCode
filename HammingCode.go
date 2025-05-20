package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Use: ./HammingCode -c/-d text.txt")

		return
	}

	file, err := os.Open(args[2])
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	if args[1] == "-c" {
		// Criptografy Mode
		Coder(file)
	} else if args[1] == "-d" {
		// Decriptografy Mode
		Decoder(file)
	} else {
		fmt.Println("Mode Unwknow")
		return
	}

}
