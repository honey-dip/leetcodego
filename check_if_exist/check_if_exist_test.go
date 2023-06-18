package check_if_exist

import "testing"

func TestCheckIfExist(t *testing.T) {
	var arr []int
	arr = append(arr, -20, 8, -6, -14, 0, -19, 14, 4)
	CheckIfExist(arr)
}
