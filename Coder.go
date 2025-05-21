package main

import (
	"fmt"
	"os"
	"strings"
)

var pow2 []byte

func CreateNewFile(file *os.File) *os.File {
	var name string
	var newName string

	name = file.Name()
	parts := strings.Split(name, ".")

	newName = parts[0] + ".dec"

	fmt.Println(newName)

	return nil
}

/*
 * Defines the set of powers from 2 to max
 */
func SetsPow2(max int) {
	var pow int

	pow = 1
	for pow <= max {
		pow2 = append(pow2, byte(pow))
		pow = pow * 2
	}
}

/*
 * Returns true if x is a power of 2
 */
func IsPow2(x int) bool {
	var size int = len(pow2)
	for i := range size {
		if x == int(pow2[i]) {
			return true
		}
	}

	return false
}

/*
 * Returns a bit array equivalent to the value of the byte
 */
func ByteToBin(bt byte) []byte {
	var i int
	var div, mod, num byte
	var binary = make([]byte, 8)

	i = 7
	num = bt
	for num > 0 {
		div = num / 2
		mod = num % 2
		num = div
		binary[i] = mod
		i--
	}

	return binary
}

func BinToBytes(bins []byte) uint32 {
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

func CalculateParityBits(hammingBlock []byte, hamBlockSize int) {
	var xor byte = 0

	for i := range hamBlockSize {
		if hammingBlock[i] == 1 {
			xor ^= byte(i + 1)
		}
	}

	xorBit := ByteToBin(xor)
	for i := range pow2 {
		hammingBlock[pow2[i]-1] = xorBit[7-i]
	}
}

func HammingCoder(byteGroup []byte) {
	var j, byteGroupSize int
	var hammingBlock = make([]byte, 0)

	/*
	 * Creates a slice of size hammingBits and fills in the elements in all fields where the index is not a power of 2.
	 */
	j = 0
	byteGroupSize = len(byteGroup)
	for i := range hammingBits {
		if !(IsPow2(i + 1)) && (j <= byteGroupSize) {
			hammingBlock = append(hammingBlock, byteGroup[j])
			j++
		} else {
			hammingBlock = append(hammingBlock, 0)
		}
	}

	CalculateParityBits(hammingBlock, hammingBits)
	bytesToWrite := BinToBytes(hammingBlock)
	fmt.Println(bytesToWrite)
}

func HammingFunc(buff []byte, size int) {
	var binByte []byte
	var binaryBlock = make([]byte, 0)

	SetsPow2(hammingBits)

	/*
	 * Unites all bytes into a block of bits
	 */
	for i := range buff {
		binByte = ByteToBin(buff[i])
		binaryBlock = append(binaryBlock, binByte...)
	}

	/*
	 * Separates the block into groups of groupSize bits to apply the Hamming code
	 */

	groupSize := dataBits
	for i := 0; i < len(binaryBlock); i += groupSize {
		end := i + groupSize
		group := binaryBlock[i:end]
		HammingCoder(group)
	}
}

func Coder(file *os.File) int {
	CreateNewFile(file)

	var numberOfBytes int = CalculateNumberOfBits() / 8
	buffer := make([]byte, numberOfBytes)

	for {
		_, err := file.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error")
			}
			break
		}
	}

	for i := range numberOfBytes {
		if i%2 == 0 {
			buffer[i] = 0b11111111
		} else {
			buffer[i] = 0b0
		}
	}

	HammingFunc(buffer, numberOfBytes)

	return 0
}
