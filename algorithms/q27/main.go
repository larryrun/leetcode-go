package main

import "github.com/larryrun/leetcode-go/utils"

/*
https://leetcode.com/problems/remove-element/
 */

func removeElement(nums []int, val int) int {
	l, r := 0, len(nums)
	for ; l < r; {
		if nums[l] == val {
			nums[l] = nums[r - 1]
			r--
		} else {
			l++
		}
	}
	return r
}

func main() {
	arr := (&[4]int{3, 2, 2, 3})[:]
	utils.AssertEq(2, removeElement(arr, 2))
	arr = (&[8]int{0, 1, 2, 2, 3, 0, 4, 2})[:]
	utils.AssertEq(5, removeElement(arr,2))
	arr = (&[1]int{2})[:]
	utils.AssertEq(1, removeElement(arr,3))
}
