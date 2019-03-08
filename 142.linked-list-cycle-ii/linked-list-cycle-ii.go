package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	var slowNew *ListNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			slowNew = head
			break
		}
	}

	if slowNew == nil { //
		return nil
	}

	for slowNew != slow {
		slowNew = slowNew.Next
		slow = slow.Next
	}
	return slowNew
}

func main() {
	
}
