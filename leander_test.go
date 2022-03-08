package leander

import "testing"

func TestValid(t *testing.T) {

	cases := map[string]struct {
		in       string
		expected bool
	}{
		"no dashes":     {"123abc", false},
		"trailing dash": {"123-abc-", false},
		"leading dash":  {"-123-abc", false},
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
