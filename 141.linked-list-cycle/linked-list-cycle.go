package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// 开始时不需要判断链表长度，因为把head赋值给fast之后，如果head为nil或者head.Next为nil，那么12行循环条件直接退出。
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {  // 这里不需要判断slow是否为空，因为slow走的路fast已经走过了。
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {

}
