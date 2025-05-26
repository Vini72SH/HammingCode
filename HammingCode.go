package main

import (
	"fmt"
	"os"
)

const dataBits byte = 247
const hammingCodeBits byte = 255

var pow2 []byte

/*
 * Defines the set of powers from 2 to max
 */
func SetsPow2(max byte) {
	var pow byte

	pow = 1
	for pow <= max && pow != 0 {
		pow2 = append(pow2, byte(pow))
		pow = pow * 2
	}

}

/*
 * Returns true if x is a power of 2
 */
func IsPow2(x byte) bool {
	var size int = len(pow2)
	for i := range size {
		if x == pow2[i] {
			return true
		}
	}

	return false
}

/*
 * Through MMC, it calculates the smallest amount of bits that are aligned
 * with the size of a byte, allowing these aligned groups of bits to be
 * read and processed by the coder and decoder
 */
func CalculateNumberOfBits() int {
	var r int
	var dBits = int(dataBits)
	var bBits = 8 // SizeOfByte

	/*
	 * Euclidean algorithm
	 */
	for bBits != 0 {
		r = dBits % bBits
		dBits = bBits
		bBits = r
	}

	return int(dataBits) * (8 / dBits)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Use: ./HammingCode -c/-d filename.txt/filename.hamming")

		return
	}

	/*
	 * Hamming code supports at most 8 bits of parity
	 */
	if (dataBits >= hammingCodeBits) || (hammingCodeBits-dataBits > 8) {
		fmt.Println("Invalid Hamming code")
		return
	}

	file, err := os.Open(args[2])
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	fmt.Println("Input file: ", file.Name())
	fmt.Println("Hamming Code Config: ", hammingCodeBits, "/", dataBits)
	SetsPow2(hammingCodeBits)
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
