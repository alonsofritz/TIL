package main

import "fmt"

func main() {
	s := "aaaaaeeeeiiioou"

	fmt.Println(charFrequency(s))
}

func charFrequency(s string) map[rune]int {
	freq := make(map[rune]int)

	for _, ch := range s {
		freq[ch]++
	}
	return freq
}
