package E_ProductOfArrayExceptSelf_238

/*
	Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].

	Example:

	Input:  [1,2,3,4]
	Output: [24,12,8,6]
	Note: Please solve it without division and in O(n).

	Follow up:
	Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)
*/

//Solution: O(1) space approach
//Algorithm
//
//1. Initialize the empty answer array where for a given index i, answer[i] would contain the product of all the numbers to the left of i.
//2. We construct the answer array the same way we constructed the L array in the previous approach. These two algorithms are exactly the
//   same except that we are trying to save up on space.
//3. The only change in this approach is that we don't explicitly build the R array from before. Instead, we simply use a variable to keep
//   track of the running product of elements to the right and we keep updating the answer array
//   by doing answer[i] = answer[i] * Ranswer[i]=answer[i]∗R. For a given index i, answer[i] contains the product of all
//   the elements to the left and R would contain product of all the elements to the right.
//   We then update R as R = R * nums[i]R=R∗nums[i]

//Complexity analysis

//Time complexity : O(N) where N represents the number of elements in the input array.
// We use one iteration to construct the array L, one to update the array answer.

//Space complexity : O(1) since don't use any additional array for our computations.
// The problem statement mentions that using the answer array doesn't add to the space complexity.

func productExceptSelf(nums []int) []int {
	// The length of the input array
	length := len(nums)

	// Final answer array to be returned
	answer := make([]int, length)

	// answer[i] contains the product of all the elements to the left
	// Note: for the element at index '0', there are no elements to the left,
	// so the answer[0] would be 1	answer[0] = 1
	answer[0] = 1
	for i := 1; i < length; i++ {
		// answer[i - 1] already contains the product of elements to the left of 'i - 1'
		// Simply multiplying it with nums[i - 1] would give the product of all
		// elements to the left of index 'i'
		answer[i] = nums[i-1] * answer[i-1]
	}

	// right contains the product of all the elements to the right
	// Note: for the element at index 'length - 1', there are no elements to the right,
	// so the R would be 1
	right := 1
	for i := length - 1; i >= 0; i-- {
		// For the index 'i', R would contain the
		// product of all elements to the right. We update R accordingly
		answer[i] *= right
		right *= nums[i]
	}
	return answer
}

func productExceptSelf3(nums []int) []int {
	product := 1
	zero := 0

	for _, num := range nums {
		if num != 0 {
			product *= num
		} else {
			zero++
		}
	}

	if zero == len(nums) {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		if zero > 1 {
			nums[i] = 0
		} else if nums[i] == 0 {
			nums[i] = product
		} else if zero == 0 {
			nums[i] = product / nums[i]
		} else {
			nums[i] = 0
		}

	}

	return nums
}

func productExceptSelf2(nums []int) []int {
	i, product, len := 0, 1, len(nums)
	result := make([]int, len)

	// result[i] = nums[0] * ... * nums[i - 1] (excludes nums[i])
	result[0] = 1
	for i = 1; i < len; i++ {
		result[i] = nums[i-1] * result[i-1]
	}

	// starting from the end of the array, since
	// result[i] = nums[0] * ... * nums[i - 1], let
	// product = nums[i + 1] * ... * nums[len - 1],
	// new result[i] = result[i] * product, update product in
	// each iteration by multiplying it with nums[i]
	product = nums[len-1]
	for i = len - 2; i >= 0; i-- {
		result[i] *= product
		product *= nums[i]
	}
	return result
}
