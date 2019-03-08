package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	//var prev *ListNode // 这样prev才是nil, 如果利用new方法创建的会带初始值
	prev := new(ListNode)
	post := cur.Next
	for cur != nil && cur.Next != nil {
		cur.Next = post.Next
		post.Next = prev.Next
		prev.Next = post
		post = cur.Next
		//cur, cur.Next, prev = cur.Next, prev, cur // 交叉赋值较为麻烦的情况下，可以使用多项式直接赋值
	}
	return prev.Next
}

// 递归方法相当于把链表的next方向转向
func reverseListRecruse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newNode = reverseListRecruse(head.Next)
	head.Next.Next = head // newNode最终相当于head.Next，这一步相当于做了个圈，即 head -> next -> head -> next ... 循环
	head.Next = nil       // 将head的next置为空，断开正向的next，只保留反向
	return newNode
}

func main() {
	link4 := &ListNode{4, nil}
	link3 := &ListNode{3, link4}
	link2 := &ListNode{2, link3}
	head := &ListNode{1, link2}
	//head.Next, head.Next.Next = head.Next.Next, nil 删除head后的结点
	//printList(head)
	printList(reverseList(head))
	//printList(reverseListRecruse(head))
}

func printList(node *ListNode) {
	for i:=0; i<=6; i++ {
		if node == nil {
			return
		}
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("\n")
}
