package src

func removeDuplicates(nums []int) int {
	if len(nums) == 1 { return 1 }

	count := 1
	for _, n := range nums[1:] {
		if nums[count-1] < n {
			nums[count] = n
			count += 1
		}
	}

	return count
}
