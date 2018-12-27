package main

import (
	"os"
	"testing"
)

func TestOverlappingSquareInches(t *testing.T) {
	fileHandle, _ := os.Open("../input.txt")
	defer fileHandle.Close()

	OverlappingSquareInches(fileHandle)
}
