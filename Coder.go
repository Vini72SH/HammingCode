package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
 * Creates an coded file with the extension .hamming
 */
func CreateCodedFile(file *os.File) *os.File {
	var name string
	var newName string

	name = file.Name()
	parts := strings.Split(name, ".")

	if strings.Compare(parts[1], "txt") != 0 {
		fmt.Println("The file isn't .txt")
		return nil
	}

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

	/*
	 * Calculates parity bits with XOR operations
	 * on indexes where the element is 1
	 */
	for i := range hammingCodeBits {
		if hammingBlock[i] == 1 {
			xor ^= byte(i + 1)
		}
	}

	/*
	 * Adds the XOR bits in reverse order to
	 * the positions whose index is a power of 2
	 */
	xorBit := ByteToBits(xor)
	for i := range pow2 {
		hammingBlock[pow2[i]-1] = xorBit[7-i]
	}
}

func HammingCoder(bitsGroup []byte, codeFile *bufio.Writer) {
	var j, bitsGroupSize int
	var hammingBlock = make([]byte, 0)

	/*
	 * Creates a slice of size hammingCodeBits and fills in the elements
	 * in all fields where the index is not a power of 2.
	 */
	j = 0
	bitsGroupSize = len(bitsGroup)
	for i := range hammingCodeBits {
		if !(IsPow2(i + 1)) && (j < bitsGroupSize) {
			hammingBlock = append(hammingBlock, bitsGroup[j])
			j++
		} else {
			hammingBlock = append(hammingBlock, 0)
		}
	}
	CalculateParityBits(hammingBlock)
	/*
	 * Writes to the encoded file one bit at a time,
	 * inserting a space at the end
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

func HammingFuncC(buff []byte, size int, codeFile *bufio.Writer) {
	var binaryByte []byte
	var binaryBlock = make([]byte, 0)

	/*
	* Unites all bytes into a block of bits
	 */
	for i := range size {
		binaryByte = ByteToBits(buff[i])
		binaryBlock = append(binaryBlock, binaryByte...)
	}

	/*
	* Separates the block into groups of groupSize bits to apply the Hamming code
	 */
	var end int
	groupSize := int(dataBits)
	blockSize := len(binaryBlock)
	for i := 0; i < blockSize; i += groupSize {
		end = min(i+groupSize, blockSize)
		group := binaryBlock[i:end]
		HammingCoder(group, codeFile)
	}

}

func Coder(file *os.File) int {
	newFile := CreateCodedFile(file)

	if newFile == nil {
		return 1
	}
	defer newFile.Close()

	var numberOfBytes int = CalculateNumberOfBits() / 8
	buffer := make([]byte, numberOfBytes)

	fmt.Println("Selected Hamming Coding")
	fmt.Println("Generating coded file...")
	/*
	 * Dijkstra probably hates me.
	 */

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(newFile)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error")
			}
			break
		}
		HammingFuncC(buffer, n, writer)
	}
	writer.Flush()

	fmt.Println("The coded file is generated:", newFile.Name())

	return 0
}
