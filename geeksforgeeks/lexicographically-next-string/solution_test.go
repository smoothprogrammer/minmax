package lexicographically_next_string_test

import (
	"testing"

	sut "github.com/minizilla/minmax/geeksforgeeks/lexicographically-next-string"
	"github.com/minizilla/testr"
)

func TestLexicographicallyNextString(t *testing.T) {
	tests := map[string]struct {
		in, out string
	}{
		"empty": {"", "a"},
		"a":     {"a", "b"},
		"z":     {"z", "za"},
		"ab":    {"ab", "ac"},
		"az":    {"az", "bz"},
		"zz":    {"zz", "zza"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			out := sut.Solution(tc.in)
			assert.Equal(out, tc.out)
		})
	}
}
