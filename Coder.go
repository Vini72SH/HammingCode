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

func HammingCoder(byteGroup []byte, hamBlockSize int) {
	var j, byteGroupSize int
	var hammingBlock = make([]byte, 0)

	/*
	 * Creates a slice of size hamBlockSize and fills in the elements in all fields where the index is not a power of 2.
	 */
	j = 0
	byteGroupSize = len(byteGroup)
	for i := range hamBlockSize {
		if !(IsPow2(i + 1)) && (j <= byteGroupSize) {
			hammingBlock = append(hammingBlock, byteGroup[j])
			j++
		} else {
			hammingBlock = append(hammingBlock, 0)
		}
	}

	fmt.Println(hammingBlock)
}

func HammingFunc(buff []byte, size int) {
	var binByte []byte
	var binaryBlock = make([]byte, 0)

	SetsPow2(31)

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
	groupSize := len(binaryBlock) / 4
	for i := 0; i < len(binaryBlock); i += groupSize {
		end := i + groupSize
		group := binaryBlock[i:end]
		HammingCoder(group, 31)
	}
}

func Coder(file *os.File) int {
	CreateNewFile(file)

	var size int = 13
	buffer := make([]byte, size)

	for {
		_, err := file.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error")
			}
			break
		}
	}

	for i := range size {
		if i%2 == 0 {
			buffer[i] = 0b11111111
		} else {
			buffer[i] = 0b0
		}
	}

	HammingFunc(buffer, size)

	return 0
}
