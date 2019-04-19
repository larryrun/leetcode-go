package main

import "github.com/larryrun/leetcode-go/utils"

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-array/
 */

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}

func main() {
	arr := (&[3]int{1, 1, 2})[:]
	utils.AssertEq(2, removeDuplicates(arr))
	arr = (&[10]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})[:]
	utils.AssertEq(5, removeDuplicates(arr))
}
