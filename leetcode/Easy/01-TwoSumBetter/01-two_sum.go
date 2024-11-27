package main

func twoSum(nums []int, target int) []int {

	for i := range len(nums) {
		first := nums[i]
		for j := range len(nums) {
			if i != j {
				second := nums[j]
				if first+second == target {
					return []int{i, j}
				}
			}
		}
	}
	return []int{}
}
