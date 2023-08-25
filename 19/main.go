package main

import "fmt"

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
		var head *Node
		dummyNode := &Node{next: head}
		temp := dummyNode
		for i := 0; i < n; i++ {
			var num int
			_, err = fmt.Scan(&num)
			if err != nil {
				return
			}
			newNode := &Node{val: num}
			temp.next = newNode
			temp = temp.next
			fmt.Printf("%d ", newNode.val)
		}
		fmt.Println()

		//show(dummyNode.next)
		show(reverse(dummyNode.next))
	}
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

func reverse(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	var pre *Node
	cur := head
	for {
		temp := cur.next
		cur.next = pre
		pre = cur
		cur = temp
		if cur == nil {
			return pre
		}
	}
}
