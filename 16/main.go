package main

import (
	"fmt"
)

func main() {
	var n int

	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		var s string
		_, err = fmt.Scanf("%s", &s)
		if err != nil {
			return
		}
		if len(s) > 50 || len(s)%2 != 0 {
			return
		}
		fmt.Println(swap(s))
	}
}

func swap(s string) string {
	var res string
	for i := 0; i < len(s); i += 2 {
		res = res + string(s[i+1]) + string(s[i])
	}
	return res
}
