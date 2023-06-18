package remove_duplicates

func removeDuplicates(nums []int) int {
	count := 1
	length := len(nums)
	for i, j := 0, 0; j < length; {
		if nums[i] == nums[j] {
			j++
		} else if nums[i] < nums[j] && j > i+1 {
			nums[i+1] = nums[j]
			i++
			count++
		} else if nums[i] < nums[j] {
			i++
			count++
		}
	}
	return count
}
