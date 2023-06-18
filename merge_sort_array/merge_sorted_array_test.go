package merge_sort_array

import "testing"

func TestMerge(t *testing.T) {
	var nums1, nums2 []int
	// nums1 = append(nums1, 1, 2, 3, 0, 0, 0)
	nums1 = append(nums1, -1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0)
	m := 5
	nums2 = append(nums2, -1, -1, 0, 0, 1, 2)
	n := 6
	Merge(nums1, m, nums2, n)
}
