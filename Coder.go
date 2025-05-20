package main

import (
	"fmt"
	"os"
)

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

func HammingFunc(buff []byte, size int) {
	var binByte []byte
	var binaryBlock = make([]byte, 0)

	for i := range buff {
		binByte = ByteToBin(buff[i])
		binaryBlock = append(binaryBlock, binByte...)
	}

	groupSize := len(binaryBlock) / 4
	for i := 0; i < len(binaryBlock); i += groupSize {
		end := i + groupSize
		group := binaryBlock[i:end]
	}
}

func Coder(file *os.File) int {
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
