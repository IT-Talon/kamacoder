package main

import (
	"errors"
	"fmt"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	list := Constructor()
	for i := 0; i < n; i++ {
		var data int
		_, err = fmt.Scan(&data)
		if err != nil {
			return
		}
		list.insert(1, data)
	}
	var m int
	_, err = fmt.Scan(&m)
	if err != nil {
		return
	}
	for i := 0; i < m; i++ {
		var s string
		_, err = fmt.Scan(&s)
		if err != nil {
			return
		}
		switch s {
		case "get":
			var index int
			_, err = fmt.Scan(&index)
			if err != nil {
				return
			}
			val, err := list.get(index)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(val)
			}
		case "delete":
			var index int
			_, err = fmt.Scan(&index)
			if err != nil {
				return
			}
			err := list.delete(index)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("delete OK")
			}
		case "insert":
			var index, val int
			_, err = fmt.Scan(&index, &val)
			if err != nil {
				return
			}
			if err = list.insert(index, val); err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("insert OK")
			}
		case "show":
			list.Show()

		}
	}
}

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	size int
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// 1 2 3
func (this *MyLinkedList) get(index int) (int, error) {
	index--
	if this.size == 0 || index < 0 || index > this.size {
		return -1, errors.New("get fail")
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val, nil
}

func (this *MyLinkedList) insert(index, val int) error {
	index--
	if index < 0 || index > this.size {
		return errors.New("insert fail")
	}
	dummyHead := &Node{next: this.head}
	cur := dummyHead
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	newNode := &Node{
		val:  val,
		next: cur.next,
	}
	cur.next = newNode
	this.size++
	this.head = dummyHead.next
	return nil
}

func (this *MyLinkedList) delete(index int) error {
	index--
	if this.size == 0 || index < 0 || index > this.size-1 {
		return errors.New("delete fail")
	}
	dummyHead := &Node{next: this.head}
	cur := dummyHead
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	this.size--
	this.head = dummyHead.next
	return nil
}

func (this *MyLinkedList) Show() {
	if this.size == 0 {
		fmt.Println("Link list is empty")
		return
	}
	cur := this.head
	for cur.next != nil {
		fmt.Printf("%d ", cur.val)
		cur = cur.next
	}
	fmt.Println(cur.val)
}
