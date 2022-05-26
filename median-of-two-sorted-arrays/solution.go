package solution

// FindMedianSortedArrays 寻找两个正序数组的中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	isEvenNum := false
	count := (len(nums1)+len(nums2))/2 + 1
	if (len(nums1)+len(nums2))%2 == 0 {
		isEvenNum = true
	}
	i, j := 0, 0
	last, curr := 0, 0
	for count > 0 {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] <= nums2[j] {
				last, curr = curr, nums1[i]
				i++
				count--
			} else {
				last, curr = curr, nums2[j]
				j++
				count--
			}
		} else if j >= len(nums2) {
			last, curr = curr, nums1[i]
			i++
			count--
		} else {
			last, curr = curr, nums2[j]
			j++
			count--
		}
	}
	if isEvenNum {
		return float64(last+curr) / 2
	} else {
		return float64(curr)
	}
}
