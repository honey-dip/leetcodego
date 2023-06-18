package sortArrayByParity

func SortArrayByParity(nums []int) []int {
	var tmp int
	for i, j := 0, 0; i < len(nums) && j < len(nums); {
		if nums[i]%2 == 1 {
			for j < len(nums) && nums[j]%2 != 0 {
				j++
			}
			if j >= len(nums) {
				break
			}
			tmp = nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i++
			j++
		} else {
			i++
			if j < i {
				j = i
			}
		}
	}
	return nums
}
