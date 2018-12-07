package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"time"
)

type event struct {
	timestamp time.Time
	text      string
}

type shift struct {
	events []event
}

type guard struct {
	shifts []shift
}

// OverlappingSquareInches yo
func OverlappingSquareInches(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)

	var list []string

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	sort.Strings(list)

	newShift := regexp.MustCompile(`\[(?P<time>)\] Guard #(?P<guard>) begins shift`)
	fallsAsleep := regexp.MustCompile(`\[(?P<time>)\] falls asleep`)
	wakesUp := regexp.MustCompile(`\[(?P<time>)\] wakes up`)

	for _, row := range list {
		switch expression {
		case condition:

		}
	}

	return list
}

func main() {
	fmt.Println(OverlappingSquareInches(os.Stdin))
}
