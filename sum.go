package main

func Sum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return nums[0] + Sum(nums[1:])
}
