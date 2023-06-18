package leetcodego

func maximumWealth(accounts [][]int) int {
	var max = 0
	var sum int
	for _, account := range accounts {
		sum = 0
		for _, v := range account {
			sum += v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
