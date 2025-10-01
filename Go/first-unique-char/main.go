package main

import "fmt"

func main() {
	s1 := "aaaaaeeeeiiioou"
	s2 := "aaaaaaeeeeeiiiiooouu"

	fmt.Println(string(firstUniqueChar(s1)), string(firstUniqueChar(s2)))
}

func firstUniqueChar(s string) rune {
	freq := charFrequency(s)
	for _, ch := range s {
		if freq[ch] == 1 {
			return ch
		}
	}
	return rune(0)
}

func charFrequency(s string) map[rune]int {
	freq := make(map[rune]int)

	for _, ch := range s {
		freq[ch]++
	}
	return freq
}
