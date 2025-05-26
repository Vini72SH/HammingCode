package main

import (
	"flag"
	"fmt"
	"os"
)

const dataBits byte = 26
const hammingCodeBits byte = 31

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
	var encodeFile string
	var decodeFile string

	flag.StringVar(&encodeFile, "c", "", "file to encode")
	flag.StringVar(&decodeFile, "d", "", "file to decode")

	flag.Parse()

	/*
	 * Hamming code supports at most 8 bits of parity
	 */
	if (dataBits >= hammingCodeBits) || (hammingCodeBits-dataBits > 8) {
		fmt.Println("Invalid Hamming code")
		return
	}

	fmt.Println("Hamming Code Config: ", hammingCodeBits, "/", dataBits)
	SetsPow2(hammingCodeBits)

	if encodeFile != "" && decodeFile != "" {
		fmt.Println("Cannot decode and encode at the same time!")
		return
	}

	if encodeFile != "" {
		file, err := os.Open(encodeFile)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			return
		}
		defer file.Close()
		Coder(file)
	} else if decodeFile != "" {
		file, err := os.Open(decodeFile)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			return
		}
		defer file.Close()
		Decoder(file)
	} else {
		fmt.Println("Unknown mode")
	}
}
