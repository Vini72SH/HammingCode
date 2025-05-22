package main

import (
	"fmt"
	"os"
	"strings"
)

/*
 * Creates an encoded file with the extension .hamming
 */
func CreateNewFile(file *os.File) *os.File {
	var name string
	var newName string

	name = file.Name()
	parts := strings.Split(name, ".")

	newName = parts[0] + ".hamming"

	newFile, err := os.Create(newName)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return nil
	}

	return newFile
}

/*
 * Returns a bit array equivalent to the value of the byte
 */
func ByteToBits(bt byte) []byte {
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

func CalculateParityBits(hammingBlock []byte) {
	var xor byte = 0

	for i := range hammingBits {
		if hammingBlock[i] == 1 {
			xor ^= byte(i + 1)
		}
	}

	xorBit := ByteToBits(xor)
	for i := range pow2 {
		hammingBlock[pow2[i]-1] = xorBit[7-i]
	}
}

func HammingCoder(byteGroup []byte, codeFile *os.File) {
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
	CalculateParityBits(hammingBlock)

	/*
	 * Writes to the encoded file one group of bits at a time, inserting a space at the end
	 */
	for i := range hammingBlock {
		var bit string
		if hammingBlock[i] == 1 {
			bit = "1"
		} else {
			bit = "0"
		}
		codeFile.WriteString(bit)
	}
	codeFile.WriteString(" ")
}

func HammingFunc(buff []byte, size int, codeFile *os.File) {
	var binaryByte []byte
	var binaryBlock = make([]byte, 0)

	/*
	* Unites all bytes into a block of bits
	 */
	for i := range buff {
		binaryByte = ByteToBits(buff[i])
		binaryBlock = append(binaryBlock, binaryByte...)
	}

	/*
	* Separates the block into groups of groupSize bits to apply the Hamming code
	 */

	groupSize := int(dataBits)
	for i := 0; i < len(binaryBlock); i += groupSize {
		end := i + groupSize
		group := binaryBlock[i:end]
		HammingCoder(group, codeFile)
	}
}

func Coder(file *os.File) int {
	newFile := CreateNewFile(file)

	if newFile == nil {
		return 1
	}
	defer newFile.Close()

	SetsPow2(byte(hammingBits))
	var numberOfBytes int = CalculateNumberOfBits() / 8
	buffer := make([]byte, numberOfBytes)

	/*
	* Dijkstra probably hates me.
	 */
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error")
			}
			break
		}
		HammingFunc(buffer, n, newFile)
	}

	return 0
}
