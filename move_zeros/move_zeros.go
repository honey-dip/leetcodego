package move_zeros

import "fmt"

func MoveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums) && j < len(nums); {
		fmt.Printf("before:i=%d, j=%d, %v\n", i, j, nums)
		if nums[i] == 0 {
			for j < len(nums) {
				if nums[j] != 0 {
					nums[i] = nums[j]
					nums[j] = 0
					break
				}
				j++
			}
		} else {
			i++
			if j < i {
				j = i
			}
		}
		fmt.Printf("after:i=%d, j=%d, %v\n", i, j, nums)
	}
}
