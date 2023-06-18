package findDisappearedNumbers

import (
	"fmt"
	"sort"
)

func FindDisappearedNumbers(nums []int) []int {
	var arr []int
	expected := 1
	sort.Sort(sort.IntSlice(nums))
	fmt.Printf("nums=%v\n", nums)
	for i, v := range nums {
		fmt.Printf("before: expected=%d, v=%d, arr=%v\n", expected, v, arr)
		if expected == v {
			expected++
		} else if expected < v {
			for expected < v {
				arr = append(arr, expected)
				expected++
			}
			if i+1 == v {
				expected++
			}
		} else if expected > v && i == len(nums)-1 {
			arr = append(arr, expected)
		}
		fmt.Printf("after: expected=%d, v=%d, arr=%v\n", expected, v, arr)
	}
	return arr
}
