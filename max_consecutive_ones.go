package leetcodego

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	max := 0
	tmp := nums[0]

	for _, v := range nums {
		if v == 1 {
			if tmp == 1 {
				count += 1
			} else {
				count = 1
			}
		} else {
			count = 0
		}
		if count > max {
			max = count
		}
		tmp = v
	}
	return max
}
