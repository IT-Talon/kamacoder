package main

import (
	"container/list"
	"strconv"
)

func main() {

	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack("hello")
	l.PushBack("world")
	e := l.Front()
	l.Remove(e)
}
func evalRPN(tokens []string) int {
	l := list.New()
	for _, token := range tokens {
		data, err := strconv.Atoi(token)
		if err != nil {
			data2 := l.Remove(l.Back())
			data1 := l.Remove(l.Back())
			l.PushBack(calculate(data1.(int), data2.(int), token))
		} else {
			l.PushBack(data)
		}
	}
	return l.Back().Value.(int)
}

func calculate(data1, data2 int, token string) int {
	switch token {
	case "+":
		return data1 + data2
	case "-":
		return data1 - data2
	case "*":
		return data1 * data2
	case "/":
		return data1 / data2
	default:
		return 0
	}
}
