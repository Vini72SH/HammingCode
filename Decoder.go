package main

import (
	"fmt"
	"os"
)

func BitsToBytes(bins []byte) uint32 {
	var pow2 uint32
	var bytes uint32

	bytes = uint32(bins[len(bins)-1])

	pow2 = 2
	for i := len(bins) - 2; i >= 0; i-- {
		bytes += uint32(bins[i]) * pow2
		pow2 = pow2 * 2
	}

	return bytes
}

func Decoder(file *os.File) int {
	fmt.Println("Descriptografia...")

	return 0
}
