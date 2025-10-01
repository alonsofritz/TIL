package main

import "fmt"

func main() {
	s := "🚀 Testando 🤔"
	rs := reverse(s)
	rsw := reverseWrong(s)

	fmt.Println(rs)
	fmt.Println(rsw)
}

func reverse(s string) string {
	rs := []rune(s)

	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func reverseWrong(s string) string {
	rs := []byte(s)

	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}

	return string(rs)
}
