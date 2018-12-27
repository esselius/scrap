package clock

import "fmt"

// Clock struct
type Clock struct {
	hour   int
	minute int
}

// New creates new Clocks
func New(hour int, minute int) (clock Clock) {
	if hour >= 0 && minute >= 0 {
		hours := 60 * hour
		minutes := hours + minute

		clock = clock.Add(minutes)
	}

	if hour < 0 {
		hours := (0 - hour) * 60
		clock = clock.Subtract(hours)
	}

	clock.hour = clock.hour % 24

	return
}

// Add minutes to Clock
func (c Clock) Add(minutes int) Clock {
	mins := (c.minute + minutes) % 60
	hours := (c.minute + minutes - mins) / 60

	c.hour = c.hour + hours
	c.minute = mins

	if c.hour == 24 {
		c.hour = 0
	}

	return c
}

// Subtract minutes from Clock
func (c Clock) Subtract(minutes int) Clock {
	mins := (c.minute + minutes) % 60
	hours := (c.minute + minutes - mins) / 60

	c.hour = c.hour - hours
	c.minute = mins

	return c
}

// String turns Clock into a string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
