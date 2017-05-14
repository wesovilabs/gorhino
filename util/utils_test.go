package util

import (
	"testing"
)

var concatStringsCases = []struct {
	input          []string
	expectedResult string
}{
	{[]string{}, ""},
	{[]string{"Home"}, "Home"},
	{[]string{"This", " ", "is", " ", "my", " ", "home"}, "This is my home"},
	{[]string{"Face", "book"}, "Facebook"},
}

func TestConcatStrings(t *testing.T) {
	var result string
	for _, scenario := range concatStringsCases {
		result = ConcatStrings(scenario.input)
		if result != scenario.expectedResult {
			t.Errorf("Strings were concated incorrectly, got: %s, expected: %s.", result, scenario.expectedResult)
		}
	}
}
