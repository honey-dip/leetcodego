package heightChecker

import (
	"fmt"
	"sort"
)

func HeightChecker(heights []int) int {
	var expected []int
	result := 0
	expected = append(expected, heights...)
	sort.Ints(expected)
	fmt.Printf("expected:=%v\n", expected)
	fmt.Printf("heights:=%v\n", heights)
	for i, _ := range heights {
		if heights[i] != expected[i] {
			result++
		}
	}
	return result
}
