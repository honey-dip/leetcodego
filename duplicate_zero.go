package leetcodego

func duplicateZeros(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			i += 2
		} else {
			i += 1
		}
	}
}
