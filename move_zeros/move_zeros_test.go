package move_zeros

import "testing"

func TestMoveZero(t *testing.T) {
	var arr []int
	arr = append(arr, 0, 1, 0, 3, 12)
	MoveZeroes(arr)
}
