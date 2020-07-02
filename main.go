package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"os"
)

func main() {
	var byteCountPtr = flag.Int("byteCount", 8, "Byte Count of final output")
	var seedPtr = flag.Int64("seed", 0, "Set random seed")
	flag.Parse()
	var path = flag.Args()[0]

	runIntegrityCheck(path, *byteCountPtr, *seedPtr)
}

func runIntegrityCheck(filePath string, byteCount int, seed int64) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stats.Size()

	var reader = bufio.NewReader(file)
	array, err := GenerateIntegrityBytes(reader, size, 1024*1024, byteCount, seed)
	if err != nil {
		panic(err)
	}

	println("File Integrity Hash:", hex.EncodeToString(array))
}
