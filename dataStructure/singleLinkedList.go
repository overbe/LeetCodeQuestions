package dataStructure

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListByArray(nums []int) *ListNode {

	pHead := &ListNode{nums[0], nil}
	pCurr := pHead

	for i, value := range nums {
		if i != 0 {
			pListNode := &ListNode{value, nil}
			pCurr.Next = pListNode
			pCurr = pListNode
		}
	}

	return pHead
}

func CreateArrayByList(l *ListNode) []int {
	list := l
	result := []int{}
	for list != nil {
		result = append(result, list.Val)
		list = list.Next
	}
	return result
}
