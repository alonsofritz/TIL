package main

func main() {

}

func rotate(nums []int, k int) []int {
	n := len(nums)
	k %= n
	return append(nums[n-k:], nums[:n-k]...)
}
