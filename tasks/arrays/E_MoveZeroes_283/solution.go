package E_MoveZeroes_283

/**
	Given an array nums, write a function to move all 0's to the end of it while maintaining the relative
	order of the non-zero elements.

Note:

1. You must do this in-place without making a copy of the array.
2. Minimize the total number of operations.

*/

func moveZeroes(nums []int) []int {

	return nums
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
