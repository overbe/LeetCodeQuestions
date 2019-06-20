package E_ReverseInteger_7

import "testing"

type testpair struct {
	input  int
	output int
}

var tests = []testpair{
	{123, 321},
	{-123, -321},
	{120, 21},
	{1534236469, 0},
}

func TestReverse(t *testing.T) {
	for _, pair := range tests {
		v := reverse(pair.input)
		if v != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", v,
			)
		}
	}
}
