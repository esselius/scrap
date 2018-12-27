package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(list []string) FreqMap {
	channel := make(chan FreqMap, len(list))

	for _, words := range list {
		go func(w string) { channel <- Frequency(w) }(words)
	}

	frequency := FreqMap{}

	for range list {
		for k, v := range <-channel {
			frequency[k] += v
		}
	}

	return frequency
}
