package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func reacts(pair string) bool {
	if pair[0] != pair[1] && strings.EqualFold(string(pair[0]), string(pair[1])) {
		return true
	}

	return false
}

func shrinkText(text string) string {
	for i := 0; i < len(text)-1; i++ {
		if reacts(string(text[i]) + string(text[i+1])) {
			return text[0:i] + text[i+2:]
		}
	}

	return text
}

func reactText(s string) int {
	meh := shrinkText(s)
	for len(meh) < len(s) {
		s, meh = meh, shrinkText(meh)
	}

	return len(meh)
}

func RecurseReact(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)

	var list []string

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	return strconv.Itoa(reactText(strings.Join(list, "")))
}

func main() {
	fmt.Println(RecurseReact(os.Stdin))
}
