/**
*	BEE 1006: https://www.beecrowd.com.br/judge/en/problems/view/1006
**/

package main

import (
	"fmt"
)

func main() {

	a_weight := 2.0
	b_weight := 3.0
	c_weight := 5.0
	var a, b, c float64

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	fmt.Printf("MEDIA = %.1f\n", media(a, b, c, a_weight, b_weight, c_weight))

}

func media(a float64, b float64, c float64, a_weight float64, b_weight float64, c_weight float64) float64 {
	return ((a * a_weight) + (b * b_weight) + (c * c_weight)) / (a_weight + b_weight + c_weight)
}
