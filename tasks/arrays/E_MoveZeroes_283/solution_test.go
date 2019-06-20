package E_MoveZeroes_283

import "testing"
import f "github.com/LeetCodeQuestions/functions"

type testpair struct {
	nums   []int
	output []int
}

var tests = []testpair{
	{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
}

func TestMoveZeroes(t *testing.T) {
	for _, pair := range tests {
		result := moveZeroes(pair.nums)
		if !f.IntArrayEquals(result, pair.output) {
			t.Error(
				"For", pair.nums,
				"expected:", pair.output,
				"got:", result,
			)
		}
	}
}
