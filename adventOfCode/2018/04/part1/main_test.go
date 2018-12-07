package main

import (
	"fmt"
	"os"
	"testing"
)

func TestOverlappingSquareInches(t *testing.T) {
	fileHandle, _ := os.Open("../input.txt")
	defer fileHandle.Close()

	fmt.Println(OverlappingSquareInches(fileHandle))
}
