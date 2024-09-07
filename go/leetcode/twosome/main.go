package main

func twoSum(nums []int, target int) []int {
	a := []int{}

	for i := 0; i < len(nums); i++ {
		for j := 1; j <= len(nums); j++ {
			if (i + j) == target {
				return append(a, 0, 1)
			} else {
				return append(a, 1, 2)
			}
		}
	}
	return a
}

func main() {
	// Test case 1:

}
