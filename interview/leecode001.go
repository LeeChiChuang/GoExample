package main

import "fmt"

// 反转链表 递归
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := ReverseNode(head.Next)
	head.Next.Next = head
	last.Next = nil

	return last
}

// 反转链表的前N个节点
var curr *ListNode
func ReverseNodeN(head *ListNode, n int) *ListNode {
	if n == 1 {
		curr = head.Next
		return head
	}

	last := ReverseNodeN(head.Next, n-1)
	head.Next.Next = head
	head.Next = curr

	return last
}

// 1 -> 2 -> 3 -> 4 -> null
// 1 -> 3 -> 2 -> 4 -> null
// 反转链表的一部分 索引去见[m, n]
func ReverseBetween(head *ListNode, m, n int) *ListNode {
	if m == 1 {
		return ReverseNodeN(head, n)
	}

	head.Next = ReverseBetween(head.Next, m-1, n-1)
	return head
}

func main()  {
	node4 := ListNode{4, nil}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	head := ListNode{1, &node2}

	n := ReverseBetween(&head, 2,3)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}