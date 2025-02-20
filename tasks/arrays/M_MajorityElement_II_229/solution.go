package M_MajorityElement_II_229

/*
	Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

	Note: The algorithm should run in linear time and in O(1) space.

	Example 1:

	Input: [3,2,3]
	Output: [3]
	Example 2:

	Input: [1,1,1,3,3,2,2,2]
	Output: [1,2]
*/

func majorityElement(nums []int) []int {
	// Since we are checking if a num appears more than 1/3 of the time
	// it is only possible to have at most 2 nums (>1/3 + >1/3 = >2/3)
	count1, count2, candidate1, candidate2 := 0, 0, 0, 1

	// Select candidates
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 <= 0 {
			// We have a bad first candidate, replace!
			candidate1, count1 = num, 1
		} else if count2 <= 0 {
			// We have a bad second candidate, replace!
			candidate2, count2 = num, 1
		} else {
			// Both candidates suck, boo!
			count1--
			count2--
		}
	}

	// Recount! Since in the first loop the count was used as
	// a weight to know if we should choose a different candidate
	// it may not be the total count. We must recount for the
	// final 2 candidates
	count1 = 0
	count2 = 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	minCount := len(nums) / 3
	result := []int{}

	if count1 > minCount && count2 > minCount {
		result = append(result, candidate1, candidate2)
	} else if count1 > minCount {
		result = append(result, candidate1)
	} else if count2 > minCount {
		result = append(result, candidate2)
	}

	return result
}

func majorityElement2(nums []int) []int {
	result := make([]int, 0)
	m := make(map[int]int)

	for _, val := range nums {
		if _, ok := m[val]; ok {
			m[val]++
		} else {
			m[val] = 1
		}
	}

	for key, value := range m {
		if value > len(nums)/3 {
			result = append(result, key)
		}
	}

	return result
}
