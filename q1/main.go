package q1

func twoSum(nums []int, target int) []int {
	numIdxMap := make(map[int]int, len(nums))
	for i, v := range nums {
		numIdxMap[v] = i
	}
	for i := range nums {
		num2 := target - nums[i]
		idx2, ok := numIdxMap[num2]
		if ok && i < idx2 {
			idxes := [2]int{i + 1, idx2 + 1}
			return idxes[:]
		}
	}
	return nil
}
