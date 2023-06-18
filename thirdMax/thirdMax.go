package thirdMax

import (
	"fmt"
	"sort"
)

func ThirdMax(nums []int) int {
	var arr []int
	for _, v := range nums {
		arr = addThird(arr, v, 3)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if len(arr) < 3 {
		return arr[0]
	} else {
		return arr[2]
	}
}

func contains(arr []int, target int) bool {
	result := false
	for _, v := range arr {
		if v == target {
			result = true
			break
		}
	}
	return result
}

func addThird(arr []int, target int, limit int) []int {
	if contains(arr, target) {
		return arr
	}
	if len(arr) < limit {
		arr = append(arr, target)
	} else {
		min := arr[0]
		minIndex := 0
		for i, v := range arr {
			if v <= min {
				min = v
				minIndex = i
			}
		}
		if min < target {
			arr[minIndex] = target
		}
	}
	fmt.Printf("arr=%v\n", arr)
	return arr
}
