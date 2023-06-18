package heightChecker

import "testing"

func TestHeightChecker(t *testing.T) {
	var arr []int
	arr = append(arr, 1, 1, 4, 2, 1, 3)
	HeightChecker(arr)
}
