package main

import "testing"

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	sum := 10
	if sum != Sum(nums) {
		t.Fatalf("wrong sum, it should not be %d, but %d", Sum(nums), sum)
	}
	t.Log("success, sum is", sum)
}
