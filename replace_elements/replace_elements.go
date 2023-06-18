package replace_elements

func replaceElements(arr []int) []int {
	biggest, tmp := 0, 0
	for i := len(arr) - 1; i >= 0; i-- {
		if i == len(arr)-1 {
			biggest = arr[i]
			arr[i] = -1
		} else {
			tmp = arr[i]
			arr[i] = biggest
			if tmp > biggest {
				biggest = tmp
			}
		}
	}
	return arr
}
