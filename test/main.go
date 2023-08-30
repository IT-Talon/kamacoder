/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import (
	"fmt"
)

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
	r := removeNthFromEnd(head, 2)
	fmt.Println(r)

}

func partition(head *ListNode, x int) *ListNode {
	dummyHead := &ListNode{Next: head}
	right := &ListNode{}
	temp := right
	cur := head
	for cur.Next != nil {
		if cur.Next.Val < x {
			cur = cur.Next
		} else {
			temp.Next = cur.Next
			temp = temp.Next
			cur.Next = cur.Next.Next
		}
	}
	temp.Next = nil
	cur.Next = right
	return dummyHead.Next
}
