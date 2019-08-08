package M_PrisonCellsAfterNDays_957

import (
	"reflect"
	"testing"
)

type testpair struct {
	nums   []int
	steps  int
	output []int
}

var tests = []testpair{
	{[]int{0, 1, 0, 1, 1, 0, 0, 1}, 7, []int{0, 0, 1, 1, 0, 0, 0, 0}},
	{[]int{0, 1, 0, 1, 1, 0, 0, 1}, 1, []int{0, 1, 1, 0, 0, 0, 0, 0}},
	{[]int{0, 1, 0, 1, 1, 0, 0, 1}, 2, []int{0, 0, 0, 0, 1, 1, 1, 0}},
}

func TestPrisonAfterNDays(t *testing.T) {
	for _, pair := range tests {
		result := prisonAfterNDays(pair.nums, pair.steps)
		if !reflect.DeepEqual(result, pair.output) {
			t.Error(
				"For", pair.nums,
				"and target", pair.steps,
				"expected:", pair.output,
				"got:", result,
			)
		}
	}
}
