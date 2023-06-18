package remove_element

func removeElement(nums []int, val int) int {
	count := len(nums)
	for i, j := 0, len(nums)-1; j >= i; {
		if nums[i] == val {
			nums[i] = -1
			nums[i] = nums[j]
			j--
			count--
		} else {
			i++
		}
	}
	return count
}
