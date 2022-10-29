package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	reqIdx := totalLen / 2
	var prevVal, curVal, idx1, idx2, virtIdx int
	for virtIdx <= reqIdx {
		prevVal = curVal
		if idx1 >= len(nums1) {
			curVal = nums2[idx2]
			idx2 += 1
		} else if idx2 >= len(nums2) || nums1[idx1] <= nums2[idx2] {
			curVal = nums1[idx1]
			idx1 += 1
		} else {
			curVal = nums2[idx2]
			idx2 += 1
		}
		virtIdx += 1
	}
	if totalLen%2 == 0 {
		return float64(prevVal+curVal) / 2
	} else {
		return float64(curVal)
	}
}
