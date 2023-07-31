/**
*	BEE 1009: https://www.beecrowd.com.br/judge/en/problems/view/1009
**/
package main

import "fmt"

type seller struct {
	name   string
	salary float64
	sales  float64
}

func (e *seller) final_salary() {

	salary := e.salary + ((e.sales * 15) / 100)

	fmt.Printf("TOTAL = R$ %.2f\n", salary)
}

func main() {

	var new_name string
	var new_salary, new_sales float64

	fmt.Scanln(&new_name)
	fmt.Scanln(&new_salary)
	fmt.Scanln(&new_sales)

	new_seller := &seller{
		name:   new_name,
		salary: new_salary,
		sales:  new_sales,
	}

	new_seller.final_salary()
}
