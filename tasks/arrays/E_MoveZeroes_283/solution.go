package E_MoveZeroes_283

/**
	Given an array nums, write a function to move all 0's to the end of it while maintaining the relative
	order of the non-zero elements.

Note:

1. You must do this in-place without making a copy of the array.
2. Minimize the total number of operations.

*/

//Solution 1 (Space Optimal, Operation Sub-Optimal) [Accepted]
//
//Complexity Analysis
//Space Complexity : O(1). Only constant space is used.
//
//Time Complexity: O(n). However, the total number of operations are still sub-optimal. The total operations (array writes) that code does is nn (Total number of elements).

func moveZeroes(nums []int) []int {
	currentNotZero := 0
	// If the current element is not 0, then we need to
	// append it just in front of last non 0 element we found.
	for i := range nums {
		if nums[i] != 0 {
			nums[currentNotZero] = nums[i]
			currentNotZero++
		}
	}
	// After we have finished processing new elements,
	// all the non-zero elements are already at beginning of array.
	// We just need to fill remaining array with 0's.
	for i := currentNotZero; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}

// Solution 2 (Optimal) [Accepted]
// Complexity Analysis
//
// Space Complexity : O(1). Only constant space is used.
//
// Time Complexity: O(n).
// However, the total number of operations are optimal. The total operations (array writes)
// that code does is Number of non-0 elements.This gives us a much better best-case (when most of the elements are 0)
// complexity than last solution. However, the worst-case (when all elements are non-0) complexity for both
// the algorithms is same.

func moveZeroes2(nums []int) []int {

	//1.All elements before the slow pointer (lastNonZeroFoundAt) are non-zeroes.
	//2. All elements between the current and slow pointer are zeroes.

	lastNonZeroFoundAt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroFoundAt], nums[i] = nums[i], nums[lastNonZeroFoundAt]
			lastNonZeroFoundAt++
		}
	}

	return nums
}
