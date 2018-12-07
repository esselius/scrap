package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMatch(t *testing.T) {
	fileHandle, _ := os.Open("../input.txt")
	defer fileHandle.Close()

	fmt.Println(Match(fileHandle))
}
