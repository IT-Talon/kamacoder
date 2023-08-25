package main

import "fmt"

func main() {
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			return
		}
		for i := 0; i < n; i++ {
			var m int
			_, err = fmt.Scan(&m)
			if err != nil {
				return
			}
			var sum int
			for i := 0; i < m; i++ {
				var t int
				_, err = fmt.Scan(&t)
				if err != nil {
					return
				}
				sum += t
			}
			fmt.Println(sum)
			if i != n-1 {
				fmt.Println()
			}
		}
	}

}
