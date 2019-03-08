package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := makeList([]int{1, 2, 3, 4, 5, 6})
	printList(head)

	head = reverseKGroup(head, 4)
	printList(head)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	start := new(ListNode)
	start.Next = head

	prev := start
	cur := head
	t := 1

	for cur != nil {
		if t%k == 0 {
			prev = reverseGroup(prev, cur.Next)
			cur = prev.Next
		} else {
			cur = cur.Next
		}
		t++
	}
	return start.Next
}

func reverseGroup(prev, post *ListNode) *ListNode {
	lprev := prev.Next
	cur := lprev.Next
	// 指针的next赋值除了用较差赋值外，在使用顺序赋值时要注意，先变的在前。如p1.next = p2.next, 再p2.next = p3.next；
	// 否则p2.next赋值为p3.next后，再p1.next = p2.next此时p2.next已经修改。
	for cur != post {
		lprev.Next = cur.Next
		cur.Next = prev.Next
		prev.Next = cur
		cur = lprev.Next
	}

	return lprev
}

func makeList(array []int) *ListNode {
	if len(array) < 1 {
		return nil
	}

	head := new(ListNode)
	head.Val = array[0]
	head.Next = nil

	cur := head
	for i := 1; i < len(array); i++ {
		newNode := new(ListNode)
		newNode.Val = array[i]
		cur.Next = newNode
		cur = cur.Next
	}
	return head
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("\n")
	return
}
