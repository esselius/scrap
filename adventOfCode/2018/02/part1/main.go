package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func numMatches(matches map[rune]int, num int) (result bool) {
	for _, v := range matches {
		if v == num {
			result = true
		}
	}
	return result
}

func matching(list []string, num int) int {
	var rows int

	for _, row := range list {
		matches := make(map[rune]int)

		for _, char := range row {
			matches[char]++
		}

		if numMatches(matches, num) {
			rows++
		}
	}

	return rows
}

// Checksum omg
func Checksum(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)

	var list []string

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	twos := matching(list, 2)
	threes := matching(list, 3)

	checksum := twos * threes

	return strconv.Itoa(checksum)
}

func main() {
	fmt.Println(Checksum(os.Stdin))
}
