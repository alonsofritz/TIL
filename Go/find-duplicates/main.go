package main

func main() {

}

func findDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	var dup []int
	for _, n := range nums {
		if seen[n] {
			dup = append(dup, n)
		}
		seen[n] = true
	}
	return dup
}
