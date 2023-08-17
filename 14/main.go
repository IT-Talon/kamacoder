package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		input := scanner.Text()
		words := strings.Fields(input)
		var res string
		for _, word := range words {
			res += word[:1]
		}
		fmt.Println(strings.ToUpper(res))
	}
}
