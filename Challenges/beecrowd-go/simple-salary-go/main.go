/**
*	BEE 1008: https://www.beecrowd.com.br/judge/en/problems/view/1008
**/
package main

import "fmt"

type employee struct {
	number       int
	hours_month  int
	payment_hour float64
}

func (e *employee) salary() {

	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", e.number, (e.payment_hour * float64(e.hours_month)))
}

func main() {

	var a, b int
	var c float64

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	new_employee := &employee{
		number:       a,
		hours_month:  b,
		payment_hour: c,
	}

	new_employee.salary()
}
