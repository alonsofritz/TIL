/**
*	BEE 1007: https://www.beecrowd.com.br/judge/en/problems/view/1007
**/

package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)

	fmt.Printf("DIFERENCA = %d", difference(a, b, c, d))
}

func difference(a int, b int, c int, d int) int {
	return ((a * b) - (c * d))
}
