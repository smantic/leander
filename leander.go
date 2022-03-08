// package leander is for working with strings following a custom convention of:
//	num-chars-num-chars...
package leander

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid takes a string input and tests that it folllows the custom convention
// we are assuming false in edgecases such as empty string
//
// estimated difficulty: low
// actual time: ~10 mins
func Valid(in string) bool {

	split := strings.Split(in, "-")

	// doesnt contain dash
	// or is otherwise empty
	if len(split) < 2 {
		return false
	}

	// main validation loop
	for i, s := range split {

		// even items should be integers.
		isEven := (i % 2) == 0

		if isEven {
			_, err := strconv.Atoi(s)
			if err != nil {
				return false
			}
		}

		// odd items should be valid ascii
		if !isEven {
			for _, c := range s {
				if !unicode.IsLetter(c) {
					return false
				}
			}
		}
	}

	return true
}
