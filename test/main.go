/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	r := reverseBetween(head, 4, 5)
	fmt.Println(r)

}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur = cur.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// 1 2 3 4
func reverse(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
		if cur == nil {
			return pre
		}
	}
	return head
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummyHead := &ListNode{Next: head}

	pre := dummyHead
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	sub := pre.Next
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		cur = cur.Next
	}
	succ := cur.Next
	cur.Next = nil
	s := reverse(sub)
	t := s
	for t.Next != nil {
		t = t.Next
	}
	t.Next = succ
	pre.Next = s
	return dummyHead.Next
}
