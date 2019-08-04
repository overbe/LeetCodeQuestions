package M_RotateList_61

import d "github.com/overbe/LeetCodeQuestions/dataStructure"

/*
Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:

Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
Example 2:

Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NULL

*/

func rotateRight(head *d.ListNode, k int) *d.ListNode {

	if head == nil || head.Next == nil && k == 1 {
		return head
	}

	currentNode := head
	//1. Finding out the length of the list
	var length = 1
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		length++
	}

	currentNode.Next = head

	j := length - (k % length)
	for i := 1; i <= j; i++ {
		currentNode = currentNode.Next
	}

	head = currentNode.Next
	currentNode.Next = nil

	return head
}
