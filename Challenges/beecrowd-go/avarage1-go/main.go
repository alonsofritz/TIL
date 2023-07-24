/**
*	BEE 1005: https://www.beecrowd.com.br/judge/en/problems/view/1005
**/

package main

import (
	"fmt"
)

func main() {

	a_weight := 3.5
	b_weight := 7.5
	var a, b float64

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	fmt.Printf("MEDIA = %.5f\n", media(a, b, a_weight, b_weight))

}

func media(a float64, b float64, a_weight float64, b_weight float64) float64 {
	return ((a * a_weight) + (b * b_weight)) / (a_weight + b_weight)
}
