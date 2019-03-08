package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	// 初始化新指针
	newA := headA
	newB := headB

	for newA != newB {
		if newA != nil {
			newA = newA.Next
		} else {
			newA = headB
		}

		if newB != nil {
			newB = newB.Next
		} else {
			newB = headA
		}
	}

	return newA
}

func main() {
	
}
