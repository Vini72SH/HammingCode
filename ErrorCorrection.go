package main

import (
	"fmt"
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

	buffer := make([]byte, 13)

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading file: ", err)
			}
			break
		}
		fmt.Printf("Read %d bytes: %q\n", n, buffer[:n])
	}

}
