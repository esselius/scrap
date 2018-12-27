package main

import (
	"os"
	"testing"
)

func TestIfReacts(t *testing.T) {
	testCases := []struct {
		desc   string
		pair   string
		result bool
	}{
		{
			desc:   "Test aA",
			pair:   "aA",
			result: true,
		},
		{
			desc:   "Test bb",
			pair:   "bb",
			result: false,
		},
		{
			desc:   "Test CC",
			pair:   "CC",
			result: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if reacts(tC.pair) != tC.result {
				t.Fail()
			}
		})
	}
}

func TestShrink(t *testing.T) {
	testCases := []struct {
		desc   string
		text   string
		result string
	}{
		{
			desc:   "Test abaABBacfF",
			text:   "abaABBacfF",
			result: "abBBacfF",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if st := shrinkText(tC.text); st != tC.result {
				t.Error("gave: " + tC.text + ", wanted: " + tC.result + ", received: " + st)
			}
		})
	}
}

func TestReactText(t *testing.T) {
	testCases := []struct {
		desc   string
		text   string
		result int
	}{
		{
			desc:   "Test abaABBacfF",
			text:   "abaABBacfF",
			result: 6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if st := reactText(tC.text); st != tC.result {
				t.Error("gave: " + tC.text + ", wanted: " + string(tC.result) + ", received: " + string(st))
			}
		})
	}
}

func TestRecurseReact(t *testing.T) {
	fileHandle, _ := os.Open("../input.txt")
	defer fileHandle.Close()

	RecurseReact(fileHandle)
}
