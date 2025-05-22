package auto_complete_feature_test

import (
	"testing"

	sut "github.com/smoothprogrammer/minmax/geeksforgeeks/auto-complete-feature"
	"github.com/smoothprogrammer/testr"
)

func TestAutoCompleteFeature(t *testing.T) {
	dict := []string{
		"i am iron man",
		"i am iron man endgame",
		"i am iron man meme",
		"i am venom",
		"i am venom sound",
		"i am venom morbius",
		"i am batman",
	}

	tests := map[string]struct {
		q   string
		res []string
	}{
		"empty":          {"", nil},
		"can't complete": {"i am man", nil},
		"incomplete": {
			"i am iro",
			[]string{
				"i am iron man",
				"i am iron man endgame",
				"i am iron man meme",
			},
		},
		"incomplete with space": {
			"i am venom ",
			[]string{
				"i am venom morbius",
				"i am venom sound",
			},
		},
		"complete": {
			"i am batman",
			[]string{
				"i am batman",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			res := sut.Solution(tc.q, dict)
			assert.Equal(res, tc.res)
		})
	}
}
