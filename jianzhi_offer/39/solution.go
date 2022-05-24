package solution

func majorityElement(nums []int) int {
	count := map[int]int{}
	mid := len(nums) / 2
	for _, v := range nums {
		if _, ok := count[v]; ok {
			count[v]++
		} else {
			count[v] = 1
		}

		if count[v] > mid {
			return v
		}
	}
	return 0
}
