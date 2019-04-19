package main

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
*/

func removeDuplicates(nums []int) int {
	nLen := len(nums)
	if nLen < 3 {
		return nLen
	}
	i := 1
	if nums[0] != nums[1] && nums[1] == nums[2] {
		i = 2
	}
	for j := 2; j < nLen; j++ {
		if nums[i] < nums[j] {
			nums[i+1] = nums[j]
			i++
			if j < nLen-1 {
				if nums[j+1] == nums[i] {
					nums[i+1] = nums[j]
					i++
					j++
				}
			} else {
				break
			}
		}
	}
	return i + 1
}
