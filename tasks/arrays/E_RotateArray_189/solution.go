package E_RotateArray_189

/*
Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Note:
	Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
	Could you do it in-place with O(1) extra space?

Solution:

Approach #4 Using Reverse [Accepted]
Algorithm

This approach is based on the fact that when we rotate the array k times, k%nk elements from the back end of the array
come to the front and the rest of the elements from the front shift backwards.
In this approach, we firstly reverse all the elements of the array. Then, reversing the first k elements followed
by reversing the rest n-kn−k elements gives us the required result.

Let n=7 and k=3

Original List                   : 1 2 3 4 5 6 7
After reversing all numbers     : 7 6 5 4 3 2 1
After reversing first k numbers : 5 6 7 4 3 2 1
After revering last n-k numbers : 5 6 7 1 2 3 4 --> Result

Complexity Analysis

Time complexity : O(n). nn elements are reversed a total of three times.

Space complexity : O(1). No extra space is used.

*/

func rotate(nums []int, k int) []int {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)

	return nums
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

/* Approach #3 Using Cyclic Replacements [Accepted]

Complexity Analysis

Time complexity : O(n). Only one pass is used.

Space complexity : O(1). Constant extra space is used.
*/

//Solution:
// func rotate(nums []int, k int) []int {
//	n := len(nums)
//	k %= n
//	count := 0
//	for start := 0; count < n; start++ {
//		current := start
//		prev := nums[start]
//		for ok := true; ok; ok = current != start {
//			next := (current + k) % n
//			temp := nums[next]
//			nums[next] = prev
//			prev = temp
//			current = next
//			count++
//		}
//
//	}
//	return nums
//}

/*
Approach #1 Brute Force [Time Limit Exceeded]
The simplest approach is to rotate all the elements of the array in k steps by rotating the elements by 1 unit in each step.

Complexity Analysis

Time complexity : O(n*k). All the numbers are shifted by one step(O(n)O(n)) k times(O(k)O(k)).
Space complexity : O(1). No extra space is used.

*/
