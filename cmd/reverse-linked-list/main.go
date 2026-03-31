package main

import "github.com/davecgh/go-spew/spew"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// list := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}
	list := &ListNode{}
	val := reverseList(list)
	spew.Dump(val)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}

// this doesnt pass all the tests
// func reverseList(head *ListNode) *ListNode {
// 	if head.Next == nil {
// 		return head
// 	}

// 	out := &ListNode{
// 		Val:  head.Val,
// 		Next: nil,
// 	}

// 	for next := head.Next; next != nil; next = next.Next {
// 		new := &ListNode{
// 			Val:  next.Val,
// 			Next: out,
// 		}
// 		out = new
// 	}

// 	return out
// }
