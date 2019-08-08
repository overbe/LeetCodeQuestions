package E_ProductOfArrayExceptSelf_238

import "testing"
import f "github.com/LeetCodeQuestions/functions"

type testpair struct {
	nums   []int
	output []int
}

var tests = []testpair{
	{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	{[]int{1, 0}, []int{0, 1}},
	{[]int{0, 0}, []int{0, 0}},
	{[]int{0, 4, 0}, []int{0, 0, 0}},
}

func TestProductExceptSelf(t *testing.T) {
	for _, pair := range tests {
		result := productExceptSelf(pair.nums)
		if !f.IntArrayEquals(result, pair.output) {
			t.Error(
				"For", pair.nums,
				"expected:", pair.output,
				"got:", result,
			)
		}
	}
}
