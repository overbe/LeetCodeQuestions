package E_RotateArray_189

import "testing"
import f "github.com/LeetCodeQuestions/functions"

type testpair struct {
	nums   []int
	steps  int
	output []int
}

var tests = []testpair{
	{[]int{1, 2, 3, 4, 5, 6, 7}, 1, []int{7, 1, 2, 3, 4, 5, 6}},
	{[]int{7, 1, 2, 3, 4, 5, 6}, 1, []int{6, 7, 1, 2, 3, 4, 5}},
	{[]int{6, 7, 1, 2, 3, 4, 5}, 1, []int{5, 6, 7, 1, 2, 3, 4}},
	{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
}

func TestRotate(t *testing.T) {
	for _, pair := range tests {
		result := rotate(pair.nums, pair.steps)
		if !f.IntArrayEquals(result, pair.output) {
			t.Error(
				"For", pair.nums,
				"and target", pair.steps,
				"expected:", pair.output,
				"got:", result,
			)
		}
	}
}
