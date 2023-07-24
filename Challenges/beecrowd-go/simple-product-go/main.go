/**
*	BEE 1004: https://www.beecrowd.com.br/judge/en/problems/view/1004
**/

package main

import (
	"fmt"
)

func main() {

	var value1, value2 int
	fmt.Scanln(&value1)
	fmt.Scanln(&value2)
	PROD := prod(value1, value2)
	fmt.Printf("PROD = %d\n", PROD)
}

func prod(v1 int, v2 int) int {
	return v1 * v2
}
