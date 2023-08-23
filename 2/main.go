package main

import "fmt"

func main() {
	for {
		var n, x, y int
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		for i := 0; i < n; i++ {
			_, err = fmt.Scan(&x, &y)
			if err != nil {
				break
			}
			fmt.Println(x + y)
		}
	}
}
