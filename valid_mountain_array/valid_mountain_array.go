package valid_mountain_array

func validMountainArray(arr []int) bool {
	result := true
	var i, j int
	for i, j = 0, len(arr)-1; i < j; {
		if arr[i+1] <= arr[i] && arr[j-1] <= arr[j] {
			result = false
			break
		}
		if arr[i+1] > arr[i] {
			i++
		}
		if arr[j-1] > arr[j] {
			j--
		}
	}
	if i == 0 || j == len(arr)-1 {
		result = false
	}
	return result
}
