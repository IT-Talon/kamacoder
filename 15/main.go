package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	for n > 0 {
		var a, b string
		_, _ = fmt.Scanln(&a)
		_, _ = fmt.Scanln(&b)
		fmt.Println(a[:len(a)/2] + b + a[len(a)/2:])
		n--
	}
}
