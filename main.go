package main

import (
	"container/list"
	"fmt"
)

func main() {
	a := make(map[string]string)
	a["a"] = "a"
	a["b"] = "b"
	delete(a, "a")
	delete(a, "a")
	fmt.Println(a)
	fmt.Println("hello")
}

func getSum(n int) int {
	var sum int
	for {
		num := n % 10
		sum += num * num
		n = n / 10
		if n == 0 {
			return sum
		}
	}
}

func isHappy(n int) bool {
	m := make(map[int]struct{})
	m[n] = struct{}{}
	sum := getSum(n)
	for {
		if sum == 1 {
			return true
		} else {
			if _, ok := m[sum]; ok {
				return false
			}
			m[sum] = struct{}{}
		}
		sum = getSum(sum)
	}
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
