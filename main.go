package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert("hello", 3))
}

// h   o
// e l
// l
// n4 0 1 2 3 2 1 0 1 2 3 2 1 0
// i n3
// 0 0   0+0
// 1 1   1+0
// 2 2   2+0
// 3 1   1+2
// 4 0    2+2
// 5 1    1+4
// 6 2
// 7 1    1+6
// 8 0
// 9 1   1+8
// 10 2
func convert(s string, numRows int) string {
	res := make([]string, numRows)
	for i, i2 := range s {
		index := i % (2*numRows - 2)
		if index >= numRows {
			index = (2*numRows - 2) - index
		}
		res[index] = res[index] + string(i2)
	}
	return strings.Join(res, "")
}
