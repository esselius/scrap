package grains

import "errors"

func Square(num int) (result uint64, err error) {
	if num <= 0 {
		return result, errors.New("")
	}

	if num > 64 {
		return result, errors.New("")
	}

	result = 1

	for i := 1; i < num; i++ {
		result = result * 2
	}

	return
}

func Total() (total uint64) {
	for i := 1; i <= 64; i++ {
		t, _ := Square(i)

		total = total + t
	}

	return
}
