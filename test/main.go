/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	cur := head
	for {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
		if cur == nil {
			return pre
		}
	}
}
