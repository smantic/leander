package leander

import (
	"math"
	"testing"
)

func TestValid(t *testing.T) {

	cases := map[string]struct {
		in       string
		expected bool
	}{
		"empty":         {"", false},
		"no dashes":     {"123abc", false},
		"trailing dash": {"123-abc-", false},
		"leading dash":  {"-123-abc", false},
		"special chars": {"123-!abc", false},
		"missing-chars": {"123-", false},
		"ok":            {"123-abc", true},
	}

	for name, test := range cases {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Valid(test.in); got != test.expected {
				t.Errorf("got: %v, expected: %v", got, test.expected)
			}
		})
	}
}

func TestAverage(t *testing.T) {

	cases := map[string]struct {
		in string

		// expected is the expected result * 1000
		expected float64
	}{
		"empty":   {"", 0},
		"invalid": {"-123", 0},
		"10":      {"10-abc-10-abc-10-abc", 10 * 1000},
		"float":   {"0-abc-0-abc-1-abc", 0.333 * 1000},
	}

	for name, test := range cases {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Average(test.in)

			// tests for 3 decimal places
			if math.Round(got*1000) != test.expected {
				t.Errorf("got: %v, expected: %v", got, test.expected)
			}
		})
	}
}

func TestWholeStory() {

	cases := map[string]struct {
		in       string
		expected string
	}{
		"empty":   {"", ""},
		"invalid": {"-123", ""},
		"ok":      {"10-abc-10-abc-10-abc", "abc abc abc"},
	}

	for name, test := range cases {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := WholeStory(test.in)

			if got != test.expected {
				t.Errorf("got: %s, expected: %s", got, test.expected)
			}
		})
	}

}
