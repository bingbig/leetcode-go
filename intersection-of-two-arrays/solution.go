package solution

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	m := map[int]struct{}{}
	for _, n := range nums1 {
		m[n] = struct{}{}
	}

	intsec := []int{}
	for _, n := range nums2 {
		if _, ok := m[n]; ok {
			intsec = append(intsec, n)
			delete(m, n)
		}
	}

	return intsec
}
