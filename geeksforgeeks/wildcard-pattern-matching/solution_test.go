package wildcard_pattern_matching_test

import (
	"testing"

	sut "github.com/minizilla/minmax/geeksforgeeks/wildcard-pattern-matching"
	"github.com/minizilla/testr"
)

func TestWildcardPatternMatching(t *testing.T) {
	tests := map[string]struct {
		text, pattern string
		matched       bool
	}{
		"no wildcard matched":   {"abcdef", "abcdef", true},
		"no wildcard unmatched": {"abcdef", "abcdez", false},
		"?: one":                {"abcdef", "?bcdef", true},
		"?: middle":             {"abcdef", "ab?def", true},
		"?: sides":              {"abcdef", "?bcde?", true},
		"?: random":             {"abcdef", "??cd?f", true},
		"?: unmatched":          {"abcdef", "???d??f", false},
		"*: full matched":       {"abcdef", "*abcdef", true},
		"*: one":                {"abcdef", "*", true},
		"*: middle":             {"abcdef", "a*f", true},
		"*: sides":              {"abcdef", "*c*", true},
		"*: random":             {"abcdef", "*cd***f", true},
		"*: multi-prefix":       {"acccdef", "*ccd**f*", true},
		"*: unmatched":          {"abcdef", "*d****d*", false},
		"combine: separate":     {"abcdef", "*d?f", true},
		"combine: unite":        {"abcdef", "*?f", true},
		"combine: unmatched":    {"abcdef", "*f?", false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			matched := sut.Solution(tc.text, tc.pattern)
			assert.Equal(matched, tc.matched)
		})
	}
}
