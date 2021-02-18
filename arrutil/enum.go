package arrutil

import (
	"strconv"
	"strings"
)

// Ints type
type Ints []int

// String to string
func (is Ints) String() string {
	ss := make([]string, len(is))
	for _, iv := range is {
		ss = append(ss, strconv.Itoa(iv))
	}

	return strings.Join(ss, ",")
}

// Has given element
func (is Ints) Has(i int) bool {
	for _, iv := range is {
		if i == iv {
			return true
		}
	}
	return false
}

// Strings type
type Strings []string

// String to string
func (ss Strings) String() string {
	return strings.Join(ss, ",")
}

// Has given element
func (ss Strings) Has(str string) bool {
	for _, s := range ss {
		if s == str {
			return true
		}
	}
	return false
}
