// package leander is for working with strings following a custom convention of:
//	num-chars-num-chars...
package leander

import (
	"math"
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

// Average takes a string following our custom conventions and outputs the
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

// WholeStory takes a string following our custom conventions and outputs all of the text parts combined.
//
// expected difficulty: low
// actual time: 7 minutes
func WholeStory(in string) string {

	if !Valid(in) {
		return ""
	}

	split := strings.Split(in, "-")
	builder := strings.Builder{}
	for i, s := range split {

		isEven := (i % 2) == 0
		if !isEven {
			builder.WriteString(s)
			builder.WriteString(" ")
		}
	}

	ret := builder.String()
	return ret[:len(ret)-1]
}

type Stats struct {
	ShortestWord      string
	LongestWord       string
	AverageWordLength float64
	AverageWords      []string
}

// StoryStats gives us some descriptive statistics of our input string following our string convention.
//
// expected difficulty: med
// actual time:
func StoryStats(in string) Stats {

	if !Valid(in) {
		return Stats{}
	}

	var (
		ret   Stats    = Stats{}
		split []string = strings.Split(in, "-")

		runingTotalWordLen int
	)

	// we do two iterations
	// 1 to find the true average
	// 2 to find all the words that match the floor/ceil of average length

	for i, s := range split {

		isEven := (i % 2) == 0
		if !isEven {

			runingTotalWordLen = runingTotalWordLen + len(s)

			// len is constant lookup
			if len(s) > len(ret.LongestWord) {
				ret.LongestWord = s
			}

			if ret.ShortestWord == "" || len(s) < len(ret.ShortestWord) {
				ret.ShortestWord = s
			}
		}

		ret.AverageWordLength = float64(runingTotalWordLen) / float64((i+1)/2)
	}

	roundedup := int(math.Ceil(ret.AverageWordLength))
	roundedDown := int(math.Floor(ret.AverageWordLength))

	for _, s := range split {
		if len(s) == roundedup || len(s) == roundedDown {
			ret.AverageWords = append(ret.AverageWords, s)
		}
	}

	return ret
}
