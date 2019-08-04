package E_MostCommonWord_819

import "testing"

type testpair struct {
	input  string
	banned []string
	output string
}

var tests = []testpair{
	{"Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}, "ball"},
	{"a, a, a, a, b,b,b,c, c", []string{"a"}, "b"},
	{"Bob. hIt, baLl", []string{"bob", "hit"}, "ball"},
	{"a.", []string{}, "a"},
	{"Bob!", []string{"hit"}, "bob"},
}

func TestReverse(t *testing.T) {
	for _, pair := range tests {
		v := mostCommonWord(pair.input, pair.banned)
		if v != pair.output {
			t.Error(
				"For", pair.input, pair.banned,
				"expected", pair.output,
				"got", v,
			)
		}
	}
}
