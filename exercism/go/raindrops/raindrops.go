package raindrops

import "strconv"

// Convert turns an int into Pling/Plang/Plong or converts int into string
func Convert(input int) (output string) {
	if input%3 == 0 {
		output = output + "Pling"
	}

	if input%5 == 0 {
		output = output + "Plang"
	}

	if input%7 == 0 {
		output = output + "Plong"
	}

	if output == "" {
		output = strconv.Itoa(input)
	}

	return
}
