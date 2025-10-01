package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 1000

	used := make(map[int]bool)
	nums := make([]int, 0, x)

	for len(nums) < x {
		num := rand.Intn(x)
		if !used[num] {
			used[num] = true
			nums = append(nums, num)
		}
	}

	target := rand.Intn(x)

	// fmt.Println("Array:", nums)
	// fmt.Println("Target:", target)

	result := twoSum(nums, target)

	fmt.Printf("twoSum (O(n)): %v, Target: %d", result, target)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, n := range nums {
		if index, exists := m[target-n]; exists {
			return []int{index, i}
		}
		m[n] = i
	}
	return nil
}

func badTwoSum(nums []int, target int) []int {
	// Abordagem O(n²)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
