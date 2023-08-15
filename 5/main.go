package main

import "fmt"

func main() {
	var a, b int
	for {
		_, _ = fmt.Scan(&a, &b)
		fmt.Printf("%d\n\n", a+b)
	}
}
