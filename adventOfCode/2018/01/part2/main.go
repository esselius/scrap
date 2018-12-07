package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func looper(list []string, lol map[int]bool, num int) (map[int]bool, int) {
	for _, text := range list {
		symbol := string(text[0])
		currentNum, _ := strconv.Atoi(string(text[1:]))

		switch symbol {
		case "+":
			num = num + currentNum
		case "-":
			num = num - currentNum
		}

		if lol[num] {
			fmt.Println("Part2: " + strconv.Itoa(num))
			os.Exit(0)
		} else {
			lol[num] = true
		}
	}

	return lol, num
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var list []string

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	num := 0
	lol := make(map[int]bool)

	lol, num = looper(list, lol, num)

	fmt.Println("Part1: " + strconv.Itoa(num))

	for {
		lol, num = looper(list, lol, num)
	}
}
