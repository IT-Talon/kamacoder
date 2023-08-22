package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	list := &Node{}
	for i := 0; i < n; i++ {
		var num int
		_, err = fmt.Scan(&num)
		if err != nil {
			return
		}
		fmt.Printf("输入了 %d", num)
		list.insert(1, num)
	}
	list.show()
}

type Node struct {
	data int
	next *Node
}

func (head *Node) show() {
	if head == nil {
		fmt.Println("Link list is empty")
		return
	}
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.data)
		cur = cur.next
	}
}

func (head *Node) get(index int) *Node {
	cur := head
	for i := 1; i < index; i++ {
		if cur == nil {
			return nil
		}
		cur = cur.next
	}
	return cur
}

// - 3 4
func (head *Node) delete(index int) {
	if head == nil {
		fmt.Println("delete fail")
		return
	}
	dummyHead := &Node{next: head}
	cur := dummyHead
	if index == 1 {
		cur.next = cur.next.next
		return
	}
	for i := 0; i < index-1; i++ {
		if cur.next.next == nil {
			fmt.Println("delete fail")
			return
		}
		cur = cur.next
	}
	cur.next = cur.next.next
	return
}

func (head *Node) insert(index, val int) {
	dummyHead := &Node{next: head}
	cur := dummyHead
	for i := 0; i < index; i++ {
		if cur.next == nil {
			fmt.Println("insert fail")
			return
		}
		cur = cur.next
	}
	//n:=Node{
	//	data: val,
	//	next: cur.next,
	//}
	cur.next = &Node{
		data: val,
		next: cur.next,
	}
	return
}
