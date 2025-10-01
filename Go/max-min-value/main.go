package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := 100

	nums := make([]int, 0, x)

	for len(nums) < x {
		num := rand.Intn(x)
		nums = append(nums, num)
	}

	fmt.Println(minMax(nums))
}

func minMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]

	for _, n := range nums {
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}
	return min, max
}
