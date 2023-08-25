package main

import (
	"fmt"
	"strings"
)

func main() {
	r := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(r)
}

func isPalindrome(s string) bool {
	var str string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			str = str + string(s[i])
		}
	}
	str = strings.ToLower(str)
	str2 := make([]string, len(str))
	for i := 0; i < len(str); i++ {
		str2[i] = string(str[len(str)-i-1])
	}
	return strings.Join(str2, "") == str
}
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
