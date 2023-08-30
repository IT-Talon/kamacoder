package main

import (
	"container/list"
	"fmt"
)

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}

type entry struct {
	key, value int
}
type LRUCache struct {
	m        map[int]*list.Element
	l        *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:        make(map[int]*list.Element),
		l:        list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if d, ok := this.m[key]; ok {
		this.l.MoveToBack(d)
		return d.Value.(entry).value
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {
	if d, ok := this.m[key]; ok {
		d.Value = entry{key: key, value: value}
		this.l.MoveToBack(d)
		return
	}
	e := this.l.PushBack(entry{key: key, value: value})
	this.m[key] = e
	if len(this.m) > this.capacity {
		rm := this.l.Front()
		this.l.Remove(rm)
		delete(this.m, rm.Value.(entry).key)
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
