package go_aspell_check

import (
	"testing"
)

type TestCase struct {
	input  string
	output string
	check  bool
}

var TestCases = []TestCase{
	{"", "", true},
	{"Hey hey heey", "        ~~~~", false},
	{"hey hey..", "", true},
	{"hey heey..", "    ~~~~  ", false},
	{"hey heey.. hey", "    ~~~~      ", false},
	{"Saves the current settings intothe", "                           ~~~~~~~", false},
	{"ISO iso Iso", "    ~~~ ~~~", false},
}

func TestNewSpeller(t *testing.T) {
	s, err := NewSpeller(map[string]string{"lang": "en_US"})
	if err != nil {
		t.Errorf("new speller should return without an error")
	}
	if s.S.Config("lang") != "en" {
		t.Errorf("")
	}
}

func TestSpeller_Check(t *testing.T) {
	speller, _ := NewSpeller(map[string]string{"lang": "en_US"})
	for _, tc := range TestCases {
		res := speller.Check(tc.input)
		if res != tc.check {
			t.Errorf("Check(\"%s\") return %t, expected: %t", tc.input, res, tc.check)
		}
	}
}

func TestSpeller_CheckWithFeedback(t *testing.T) {
	speller, _ := NewSpeller(map[string]string{"lang": "en_US"})
	for _, tc := range TestCases {
		res := speller.CheckWithFeedback(tc.input)
		if res != tc.output {
			t.Errorf("\nCheckWithFeedback(\"%s\")\n"+
				"return            \"%s\",\n"+
				"expected:         \"%s\"", tc.input, res, tc.output)
		}
	}
}
