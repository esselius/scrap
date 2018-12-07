package main

import (
	"fmt"
	"os"
	"testing"
)

func TestChecksum(t *testing.T) {
	fileHandle, _ := os.Open("../input.txt")
	defer fileHandle.Close()

	fmt.Println(Checksum(fileHandle))
}
