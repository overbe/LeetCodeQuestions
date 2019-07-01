package E_MajorityElement_169

import "sort"

/*
	Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

	You may assume that the array is non-empty and the majority element always exist in the array.

	Example 1:

	Input: [3,2,3]
	Output: 3
	Example 2:

	Input: [2,2,1,1,1,2,2]
	Output: 2
*/

//Solution: Boyer-Moore Voting Algorithm
//Complexity Analysis
//
//Time complexity : O(n)
//Boyer-Moore performs constant work exactly nn times, so the algorithm runs in linear time.
//
//Space complexity : O(1)
//Boyer-Moore allocates only constant additional memory.

func majorityElement(nums []int) int {

	var majority int
	count := 0

	for _, value := range nums {
		if count == 0 {
			majority = value
			count++
		} else if value == majority {
			count++
		} else {
			count--
		}
	}

	return majority
}

//Solution 2: Sorting

//Algorithm
//For this algorithm, we simply do exactly what is described: sort nums, and return the element in question.
//For each example, the line below the array denotes the range of indices that are covered by a majority element that
// happens to be the array minimum. As you might expect, the line above the array is similar, but for the
// case where the majority element is also the array maximum.
//Complexity Analysis
//
//Time complexity : O(nlgn)
//Sorting the array costs O(nlgn)time in Python and Java, so it dominates the overall runtime.
//
//Space complexity : O(1) or (O(n))
//We sorted nums in place here - if that is not allowed, then we must spend linear additional space on a copy of nums and sort the copy instead.

func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
