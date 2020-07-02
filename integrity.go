package main

import (
	"bufio"
	"math/bits"
)

func GenerateIntegrityBytes(reader *bufio.Reader, fileSize int64, chunkSize int64, byteCount int, seed int64) ([]byte, error) {
	var byteGen = NewRandByteGen(seed)

	var bitCount = byteCount * 8
	var xorBitArray = make([]bool, bitCount)

	for fileSize != 0 {
		var currentChunkSize = chunkSize
		if chunkSize > fileSize {
			currentChunkSize = fileSize
		}

		var chunkBytes = make([]byte, currentChunkSize)

		_, err := reader.Read(chunkBytes)
		if err != nil {
			return nil, err
		}
		for _, currentByte := range chunkBytes {
			for i := 0; i < bitCount; i++ {
				xorBitArray[i] = xorBitArray[i] != (bits.OnesCount8(currentByte&byteGen.RandByte())&1 == 1)
			}
		}

		fileSize -= currentChunkSize
	}

	xorByteArray := make([]byte, byteCount)
	for currentByteCount := 0; currentByteCount < byteCount; currentByteCount++ {
		var currentByte byte = 0
		for bitCount := 0; bitCount < 8; bitCount++ {
			if xorBitArray[currentByteCount*8+bitCount] {
				currentByte += 1
			}
			currentByte <<= 1
		}
		xorByteArray[currentByteCount] = currentByte
	}

	return xorByteArray, nil
}
