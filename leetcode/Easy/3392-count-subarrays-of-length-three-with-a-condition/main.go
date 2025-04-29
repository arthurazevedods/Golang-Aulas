package main

func countSubarrays(nums []int) (result int) {
	for i := 1; i+1 < len(nums); i++ {
		if (nums[i-1]+nums[i+1])*2 == nums[i] {
			result++
		}
	}
	return
}
