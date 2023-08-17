package main

import "fmt"

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		if n == 0 {
			break
		}
		var sum, i int
		for n > 0 {
			_, err = fmt.Scan(&i)
			if err != nil {
				return
			}
			sum += i
			n--
		}
		fmt.Println(sum)
	}

}
