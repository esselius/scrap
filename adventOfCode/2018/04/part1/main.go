package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type sleep struct {
	start time.Time
	end   time.Time
}

func (s sleep) duration() time.Duration {
	return s.end.Sub(s.start)
}

type guard struct {
	id     int
	sleeps []sleep
}

func (g guard) totalSleep() (totalSleep int) {
	for _, s := range g.sleeps {
		totalSleep = totalSleep + int(s.duration().Minutes())
	}

	return
}

// OverlappingSquareInches yo
func OverlappingSquareInches(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)

	var list []string

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	sort.Strings(list)

	newShift := regexp.MustCompile(`\[(?P<time>.*)\] Guard #(?P<guard>\d+) begins shift`)
	fallsAsleep := regexp.MustCompile(`\[(?P<time>.*)\] falls asleep`)
	wakesUp := regexp.MustCompile(`\[(?P<time>.*)\] wakes up`)

	var id int
	var sleepz sleep
	guards := make(map[int]guard)

	for _, row := range list {
		switch {
		case newShift.MatchString(row):
			match := newShift.FindStringSubmatch(row)
			id, _ = strconv.Atoi(match[2])
		case fallsAsleep.MatchString(row):
			match := fallsAsleep.FindStringSubmatch(row)
			start, _ := time.Parse("2006-01-02 15:04", match[1])

			sleepz = sleep{
				start: start,
			}
		case wakesUp.MatchString(row):
			match := wakesUp.FindStringSubmatch(row)
			end, _ := time.Parse("2006-01-02 15:04", match[1])
			sleepz.end = end

			guard := guards[id]
			guard.id = id
			guard.sleeps = append(guard.sleeps, sleepz)

			guards[id] = guard
		}
	}

	sleepyGuard := guard{}

	for _, guard := range guards {
		if guard.totalSleep() > sleepyGuard.totalSleep() {
			sleepyGuard = guard
		}
	}

	fmt.Println("Sleepy guard: " + strconv.Itoa(sleepyGuard.id))

	mins := make([]int, 60)

	for _, s := range sleepyGuard.sleeps {
		for m := s.start.Minute(); m <= s.end.Minute(); m++ {
			mins[m]++
		}
	}

	var maxMin int
	for i, min := range mins {
		if mins[maxMin] < min {
			maxMin = i
		}
	}

	fmt.Println("Sleepy time: 00:" + strconv.Itoa(maxMin))
	fmt.Println("Result: " + strconv.Itoa(sleepyGuard.id*maxMin))

	return list
}

func main() {
	fmt.Println(OverlappingSquareInches(os.Stdin))
}
