package main

import (
	"fmt"
)

func main() {
	isHappy(3)
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
