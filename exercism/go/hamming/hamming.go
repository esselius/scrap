package hamming

import "errors"

// Distance calculates the Hamming difference between two DNA strands
func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		return -1, errors.New("")
	}

	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
