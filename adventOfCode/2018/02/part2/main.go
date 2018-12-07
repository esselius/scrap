package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func gimmieMatch(list []string) (result []string) {
	for i, row := range list {
		for _, rowz := range append(list[:i], list[i+1:]...) {
			var matches int
			length := len(rowz)

			for j := range row {
				if row[j] == rowz[j] {
					matches++
				}
			}
			if matches == length-1 {
				result = append(result, rowz, row)
			}
		}
	}

	return result
}

func answer(list []string) (result string) {
	for i, a := range list[0] {
		if byte(a) != list[1][i] {
			result = list[0][:i] + list[0][i+1:]
			break
		}
	}

	return result
}

// Match yo
func Match(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)

	var list []string

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	return answer(gimmieMatch(list))
}

func main() {
	fmt.Println(Match(os.Stdin))
}
