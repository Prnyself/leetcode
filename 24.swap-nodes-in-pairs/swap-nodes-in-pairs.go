package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode) *ListNode {
	start := new(ListNode)
	start.Next = head

	cur := start
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
	//for i := 0; i < 3; i++ {
		cur.Next, cur.Next.Next, cur.Next.Next.Next = cur.Next.Next, cur.Next.Next.Next, cur.Next
		cur = cur.Next.Next
	}
	return start.Next
}

func main() {
	//link6 := &ListNode{6, nil}
	link5 := &ListNode{5, nil}
	link4 := &ListNode{4, link5}
	link3 := &ListNode{3, link4}
	link2 := &ListNode{2, link3}
	head := &ListNode{1, link2}
	//head.Next, head.Next.Next = head.Next.Next, nil 删除head后的结点
	//printList(head)
	printList(swapNodes(head))
}

func printList(node *ListNode) {
	for i := 0; i <= 10; i++ {
		if node == nil {
			return
		}
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("\n")
}
