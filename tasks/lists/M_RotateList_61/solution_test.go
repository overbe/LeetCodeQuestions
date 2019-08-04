package M_RotateList_61

import (
	"testing"
)
import f "github.com/overbe/LeetCodeQuestions/functions"
import d "github.com/overbe/LeetCodeQuestions/dataStructure"

type testpair struct {
	input  []int
	steps  int
	output []int
}

var tests = []testpair{
	{[]int{1, 2, 3, 4, 5}, 1, []int{5, 1, 2, 3, 4}},
	{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
	{[]int{0, 1, 2}, 4, []int{2, 0, 1}},
}

func TestRotateRight(t *testing.T) {

	for _, pair := range tests {

		inputList := d.CreateListByArray(pair.input)
		inputList = rotateRight(inputList, pair.steps)
		result := d.CreateArrayByList(inputList)

		if !f.IntArrayEquals(result, pair.output) {
			t.Error(
				"For", pair.input,
				"and target", pair.steps,
				"expected:", pair.output,
				"got:", result,
			)
		}
	}
}
