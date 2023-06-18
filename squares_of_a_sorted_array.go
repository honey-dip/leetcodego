package leetcodego

func absInt(num int) int {
	if num > 0 {
		return num
	}
	return 0 - num
}

func sortedSquares(nums []int) []int {
	fore := 0
	back := len(nums) - 1

	var answer []int

	for fore != back {
		if absInt(nums[fore]) < absInt(nums[back]) {
			back -= 1
		} else {
			fore += 1
		}
	}

	for fore >= 0 || back <= len(nums)-1 {
		if fore < 0 {
			answer = append(answer, nums[back]*nums[back])
			back += 1
		} else if back > len(nums)-1 {
			answer = append(answer, nums[fore]*nums[fore])
			fore -= 1
		} else if fore == back {
			answer = append(answer, nums[fore]*nums[fore])
			fore -= 1
			back += 1
		} else if absInt(nums[back]) >= absInt(nums[fore]) {
			answer = append(answer, nums[fore]*nums[fore])
			fore -= 1
		} else {
			answer = append(answer, nums[back]*nums[back])
			back += 1
		}
	}
	return answer
}
