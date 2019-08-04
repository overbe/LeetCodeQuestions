package E_MajorityElement_169

import "testing"

type testpair struct {
	nums   []int
	output int
}

var tests = []testpair{
	{[]int{3, 2, 3}, 3},
	{[]int{2, 2, 1, 1, 1, 1, 1, 2, 2}, 1},
}

func TestMoveZeroes(t *testing.T) {
	for _, pair := range tests {
		result := majorityElement(pair.nums)
		if result != pair.output {
			t.Error(
				"For", pair.nums,
				"expected:", pair.output,
				"got:", result,
			)
		}
	}
}
