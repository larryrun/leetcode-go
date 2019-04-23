package q189

/*
https://leetcode.com/problems/rotate-array/
 */

func rotate(nums []int, k int)  {
	if k == 0 {
		return
	}
	nLen := len(nums)
	if k > nLen {
		k = k%nLen
	}
	reverse(nums[nLen-k:])
	reverse(nums[:nLen-k])
	reverse(nums)
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for ;l<r; {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
