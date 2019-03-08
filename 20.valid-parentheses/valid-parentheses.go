package main

import "fmt"

type charStack struct {
	Val int32
	Next *charStack
}

func pop(head *charStack) (stack *charStack, poped *charStack) {
	if head == nil {
		return nil, nil
	}
	poped = head
	stack = head.Next
	return
}

func push(head *charStack, node *charStack) *charStack {
	node.Next = head
	head = node
	return head
}

func main() {
	s := "{[]}"
	fmt.Println(isValid(s))
	//var head *charStack
	//node := &charStack{Val:41, Next:nil}
	//head = push(head, node)
	//node = &charStack{Val: 91, Next: nil}
	//head = push(head, node)
	//printList(head)
	//
	//var poped *charStack
	//head, poped = pop(head)
	//fmt.Println(poped)
	//printList(head)
}

func isPair(c1, c2 int32) bool {
	return (c1 == '(' && c2 == ')') || (c1 == '[' && c2 == ']') || (c1 == '{' && c2 == '}')
}

func isLeft(c int32) bool {
	return (c == '(') || (c == '[') || (c == '{')
}

func isRight(c int32) bool {
	return (c == ')') || (c == ']') || (c == '}')
}

func isValid(s string) bool {
	var head *charStack
	for _, v := range s {
		fmt.Println(v)
		if isLeft(v) {
			fmt.Println("left:", v)
			node := &charStack{Val: v, Next: nil}
			head = push(head, node)
			printList(head)
			fmt.Println("left push:", v)
		} else {
			fmt.Println("right:", v)
			if head == nil {
				return false
			}
			var node *charStack
			head, node = pop(head)
			fmt.Println(node)
			if node == nil || !isPair(node.Val, v) {
				return false
			}
		}
	}
	return head == nil
}

func printList(node *charStack) {
	for i:=0; i<=6; i++ {
		if node == nil {
			return
		}
		fmt.Printf("%c -> ", node.Val)
		node = node.Next
	}
	fmt.Printf("\n")
}