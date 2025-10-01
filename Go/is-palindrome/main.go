package main

import "fmt"

func main() {
	s := "natan"
	palindrome := isPalindrome(s)

	fmt.Println(palindrome)
}

func isPalindrome(s string) bool {
	rs := []rune(s)

	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		if rs[i] != rs[j] {
			return false
		}
	}
	return true
}
