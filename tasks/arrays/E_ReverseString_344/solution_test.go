package E_ReverseString_344

import (
	"reflect"
	"testing"
)

type testpair struct {
	input  []byte
	output []byte
}

var tests = []testpair{
	{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
	{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
}

func TestRotate(t *testing.T) {
	for _, pair := range tests {
		result := reverseString(pair.input)
		if !reflect.DeepEqual(result, pair.output) {
			t.Error(
				"For", pair.input,
				"expected:", pair.output,
				"got:", result,
			)
		}
	}
}
