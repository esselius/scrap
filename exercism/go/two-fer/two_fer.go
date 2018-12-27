package twofer

// ShareWith takes a parameter and concats strings
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
