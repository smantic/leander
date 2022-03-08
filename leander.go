// package leander is for working with strings following a custom convention of:
//	num-chars-num-chars...
package leander

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid takes a string input and tests that it folllows the custom convention.
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
			if len(s) == 0 {
				return false
			}
			for _, c := range s {
				if !unicode.IsLetter(c) {
					return false
				}
			}
		}
	}

	return true
}

// Average takes a string following our custom convents and outputs the
// average of all the numbers contained in the string.
//
// estimated difficulty: low
// actual time: 15 mins
func Average(in string) float64 {

	if !Valid(in) {
		return 0
	}

	var (
		total = 0
		freq  = 0
		split = strings.Split(in, "-")
	)

	for i, s := range split {

		// even items should be integers.
		isEven := (i % 2) == 0

		if isEven {
			num, err := strconv.Atoi(s)
			if err != nil {
				// should be non-reachable
				return 0
			}

			total = total + num
			freq = freq + 1
		}
	}

	return float64(total) / float64(freq)
}
