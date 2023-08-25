package main

import (
	"fmt"
)

func main() {
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}
		if n == 0 {
			fmt.Println("list is empty")
			continue
		}

		dummyHead := &Node{}
		cur := dummyHead
		for n > 0 {
			var val int
			_, err = fmt.Scan(&val)
			if err != nil {
				return
			}
			node := &Node{val: val}
			cur.next = node
			cur = cur.next
			n--
		}
		show(dummyHead.next)
		show(removeDuplicates(dummyHead.next))
	}

}

func removeDuplicates(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	pre := head
	for cur.next != nil {
		if pre.val == cur.next.val {
			cur.next = cur.next.next
		} else {
			cur = cur.next
			pre = pre.next
		}
	}
	return head
}

func show(head *Node) {
	if head == nil {
		return
	}
	cur := head
	for {
		fmt.Printf("%d ", cur.val)
		cur = cur.next
		if cur == nil {
			fmt.Println()
			return
		}
	}
}

type Node struct {
	val  int
	next *Node
}
