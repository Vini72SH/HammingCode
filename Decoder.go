package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
 * Creates an decoded file with the extension .dec
 */
func CreateDecodedFile(file *os.File) *os.File {
	var name string
	var newName string

	name = file.Name()
	parts := strings.Split(name, ".")

	if strings.Compare(parts[1], "hamming") != 0 {
		fmt.Println("The file isn't .hamming")
		return nil
	}

	newName = parts[0] + ".dec"

	newFile, err := os.Create(newName)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return nil
	}

	return newFile
}

func BitsToByte(bits []byte) byte {
	var i int
	var pow byte
	var finalByte byte

	i = 7
	pow = 1
	finalByte = 0
	for i >= 0 {
		finalByte += pow * bits[i]
		pow = pow * 2
		i--
	}

	return finalByte

}

func HammingDecoder(byteGroup []byte) {
	var xor byte = 0

	for i := range byteGroup {
		if byteGroup[i] == 1 {
			xor ^= byte(i + 1)
		}
	}

	if xor != 0 {
		byteGroup[xor-1] ^= 1
	}
}

func HammingFuncD(buff []byte, size int, decodeFile *bufio.Writer) {
	var end int
	var binaryBlock []byte = make([]byte, 0)
	var hammingBlock []byte = make([]byte, 0)

	/*
	 * Unifying block bits and removing spaces
	 */
	for i := range size {
		if buff[i] == 49 {
			hammingBlock = append(hammingBlock, 1)
		} else if buff[i] == 48 {
			hammingBlock = append(hammingBlock, 0)
		}
	}

	/*
	 * Separates the block into hammingCodeBits size groups
	 * and performs error correction
	 */
	groupSize := int(hammingCodeBits)
	blockSize := len(hammingBlock)
	for i := 0; i < blockSize; i += groupSize {
		end = min(i+groupSize, blockSize)
		group := hammingBlock[i:end]
		HammingDecoder(group)
		/*
		 * After correcting the errors, it removes
		 * the parity bits in a new bit block
		 */
		for j := 0; j < len(group); j++ {
			if !(IsPow2(byte(j + 1))) {
				binaryBlock = append(binaryBlock, group[j])
			}
		}
	}

	var totalBytes = len(binaryBlock) / 8
	var bytesToWrite = make([]byte, totalBytes)

	/*
	 * Converts the block of bits into a array of bytes
	 * to write to the file
	 */
	var n int = 0
	for i := range totalBytes {
		start := i * 8
		end := start + 8
		currentByte := BitsToByte(binaryBlock[start:end])
		if currentByte != 0 {
			bytesToWrite[n] = currentByte
			n++
		}
	}

	decodeFile.Write(bytesToWrite[0:n])
}

func Decoder(file *os.File) int {
	newFile := CreateDecodedFile(file)

	if newFile == nil {
		return 1
	}
	defer newFile.Close()

	/*
	 * Based on the number of bits needed to align with the bytes,
	 * it defines the number of bytes it will read from the coded file,
	 * including one byte to read the space inserted by the coder
	 */
	var numberOfGroups = CalculateNumberOfBits() / int(dataBits)
	var numberOfBytes = numberOfGroups * (1 + int(hammingCodeBits))
	buffer := make([]byte, numberOfBytes)

	fmt.Println("Selected Hamming Decoding")
	fmt.Println("Generating decoded file...")
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
		HammingFuncD(buffer, n, writer)
	}
	writer.Flush()

	fmt.Println("The decoded file is generated:", newFile.Name())

	return 0
}
