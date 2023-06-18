package findDisappearedNumbers

import "testing"

func TestFindDisappearedNumbers(t *testing.T) {
	var arr []int
	arr = append(arr, 1, 1, 2, 2)
	FindDisappearedNumbers(arr)
}
