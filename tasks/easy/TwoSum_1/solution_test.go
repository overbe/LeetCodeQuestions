package TwoSum_1

import "testing"
import f "github.com/LeetCodeQuestions/tasks/functions"

type testpair struct {
	nums   []int
	target int
	output []int
}

var tests = []testpair{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{2, 7, 11, 15, 9}, 13, []int{0, 2}},
}

func TestTwoSum(t *testing.T) {
	for _, pair := range tests {
		result := twoSum(pair.nums, pair.target)
		if !f.IntArrayEquals(result, pair.output) {
			t.Error(
				"For", pair.nums,
				"and target", pair.target,
				"expected:", pair.output,
				"got:", result,
			)
		}
	}
}
