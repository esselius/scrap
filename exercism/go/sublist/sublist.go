package sublist

type Relation string

const (
	equal     Relation = "equal"
	unequal   Relation = "unequal"
	sublist   Relation = "sublist"
	superlist Relation = "superlist"
)

func emptyLists(l1, l2 []int) bool {
	return len(l1) == 0 && len(l2) == 0
}

func isEqual(l1, l2 []int) bool {
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}

	return true
}

func isSublist(l1, l2 []int) bool {
	if len(l1) == 0 {
		return true
	}

	if len(l1) == len(l2) {
		return false
	}

	for i := 0; i <= len(l2)-len(l1); i++ {
		if isEqual(l1, l2[i:]) {
			return true
		}
	}

	return false
}

func isSuperlist(l1, l2 []int) bool {
	if len(l2) == 0 {
		return true
	}

	if len(l1) == len(l2) {
		return false
	}

	for i := 0; i <= len(l1)-len(l2); i++ {
		if isEqual(l2, l1[i:]) {
			return true
		}
	}

	return false
}

func Sublist(l1, l2 []int) Relation {
	switch {
	case emptyLists(l1, l2):
		return equal
	case isSublist(l1, l2):
		return sublist
	case isSuperlist(l1, l2):
		return superlist
	case isEqual(l1, l2):
		return equal
	}

	return unequal
}
