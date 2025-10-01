package main

import "fmt"

func main() {
	s1 := "amor"
	s2 := "roma"
	fmt.Println(isAnagrama(s1, s2))
}

func isAnagrama(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	freq := make(map[rune]int)
	for _, ch := range s1 {
		freq[ch]++
	}
	for _, ch := range s2 {
		freq[ch]--
		if freq[ch] < 0 {
			return false
		}
	}
	return true
}
