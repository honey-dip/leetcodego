package leetcodego

func findNumbers(nums []int) int {
	var digit, comp int
	answer := 0
	for _, num := range nums {
		digit = 0
		comp = 1
		for i := 0; num >= comp; i++ {
			digit += 1
			comp = comp * 10
		}
		if digit%2 == 0 {
			answer += 1
		}
	}
	return answer
}
