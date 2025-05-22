package main

import (
	"fmt"
	"os"
)

const dataBits byte = 26
const hammingBits byte = 31

func CalculateNumberOfBits() int {
	var r int
	var dBits = int(dataBits)
	var hBits = 8 // SizeOfByte

	for hBits != 0 {
		r = dBits % hBits
		dBits = hBits
		hBits = r
	}

	return int(dataBits) * (8 / dBits)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Use: ./HammingCode -c/-d text.txt")

		return
	}

	if (dataBits >= hammingBits) || (hammingBits-dataBits > 8) {
		fmt.Println("Invalid Hamming code")
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
