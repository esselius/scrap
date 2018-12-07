package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type cut struct {
	id       int
	topEdge  int
	leftEdge int
	width    int
	height   int
}

// OverlappingSquareInches yo
func OverlappingSquareInches(reader io.Reader) (overlap int) {
	scanner := bufio.NewScanner(reader)

	var list []string

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	re := regexp.MustCompile(`#(?P<id>\d+) @ (?P<leftEdge>\d+),(?P<topEdge>\d+): (?P<width>\d+)x(?P<height>\d+)`)

	cuts := make([]cut, len(list))

	for i, row := range list {
		match := re.FindStringSubmatch(row)

		id, _ := strconv.Atoi(match[1])
		leftEdge, _ := strconv.Atoi(match[2])
		topEdge, _ := strconv.Atoi(match[3])
		width, _ := strconv.Atoi(match[4])
		height, _ := strconv.Atoi(match[5])

		cuts[i] = cut{
			id:       id,
			leftEdge: leftEdge,
			topEdge:  topEdge,
			width:    width,
			height:   height,
		}
	}

	square := [1000][1000]int{}

	for _, c := range cuts {
		for i := c.leftEdge; i < c.leftEdge+c.width; i++ {
			for j := c.topEdge; j < c.topEdge+c.height; j++ {
				square[i][j]++
			}
		}
	}

	for _, c := range cuts {
		var overomg int

		for i := c.leftEdge; i < c.leftEdge+c.width; i++ {
			for j := c.topEdge; j < c.topEdge+c.height; j++ {
				if square[i][j] > 1 {
					overomg++
				}
			}
		}

		if overomg == 0 {
			overlap = c.id
		}
	}

	return overlap
}

func main() {
	fmt.Println(OverlappingSquareInches(os.Stdin))
}
