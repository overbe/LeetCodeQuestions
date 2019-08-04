package M_MajorityElement_II_229

import "testing"
import f "github.com/LeetCodeQuestions/functions"

type testpair struct {
	nums   []int
	output []int
}

var tests = []testpair{
	{[]int{3, 2, 3}, []int{3}},
	{[]int{1, 1, 1, 3, 3, 2, 2, 2}, []int{1, 2}},
}

func TestMoveZeroes(t *testing.T) {
	for _, pair := range tests {
		result := majorityElement(pair.nums)
		if !f.IntArrayEquals(result, pair.output) {
			t.Error(
				"For", pair.nums,
				"expected:", pair.output,
				"got:", result,
			)
		}
	}
}
