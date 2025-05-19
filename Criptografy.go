package main

import (
	"fmt"
	"os"
)

const max byte = 0b11111111

func GetFirstNBits(bt byte, n byte) byte {
	if n > 8 {
		return 0
	}

	var lastBits byte = 8 - n
	var mask byte = max << lastBits

	return bt & mask
}

func GetLastNBits(bt byte, n byte) byte {
	if n > 8 {
		return 0
	}

	var mask byte = (1 << n) - 1

	return bt & mask
}

func HammingFunc(buff []byte, size int) {
	var i, j, diff byte
	var currentBits uint8
	var byteShift, currentByte byte
	remainingBits := make([]uint8, size)
	var newBytes [4]uint32

	// Defines that all bits of the bytes are available
	for i := range size {
		remainingBits[i] = 8
	}

	i = 0
	j = 0
	currentBits = 0
	byteShift = 24
	for i < 4 {
		/*
		 * If the current uint has already received 26 bits of input, then go to the next group of 26 bits
		 */
		if currentBits == 26 {
			/*
			 * Each group of 26 bits creates a gap of 2 bits for the next group, (i + 1) * 2 cancels this out by guaranteeing a shift that covers this gap
			 */
			byteShift = 24 + 2*byte(i+1)
			currentBits = 0
			newBytes[i] = newBytes[i] >> 6
			i++
			continue
		}

		/*
		 * If there are no more bits left in byte j, move on to the next byte
		 */
		if remainingBits[j] == 0 {
			j++
			continue
		}

		/*
		 * The entire byte is free for use
		 */
		if remainingBits[j] == 8 {
			if (currentBits + 8) <= 26 {
				/*
				 * If using the byte does not exceed the maximum group size of 26 bits, then we can use it
				 */
				currentByte = GetFirstNBits(buff[j], 8)
				currentBits += 8
				remainingBits[j] = 0
			} else {
				/*
				 * If not, we should break the byte and only take the first few bits
				 */
				diff = 26 - currentBits
				currentByte = GetFirstNBits(buff[j], diff)
				currentBits += diff
				remainingBits[j] -= diff
			}
		} else {
			/*
			 * If the integer byte cannot be used, then it has already been broken and we must take the remaining bits of the byte
			 */
			currentByte = GetLastNBits(buff[j], remainingBits[j])
			currentBits += remainingBits[j]
			remainingBits[j] = 0
		}

		/*
		 * Add the current byte to the uint by shifting the bits to the correct positions.
		 */
		newBytes[i] |= uint32(currentByte) << uint32(byteShift)
		byteShift -= 8
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("%d | %b\n", i, newBytes[i])
	}
}

func Cripto(file *os.File) int {
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

	for i := 0; i < size; i++ {
		if i%2 == 0 {
			buffer[i] = 0b11111111
		} else {
			buffer[i] = 0b0
		}
	}

	HammingFunc(buffer, size)

	return 0
}
