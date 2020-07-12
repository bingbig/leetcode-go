package solution

func smallerNumbersThanCurrent(nums []int) []int {
	l := len(nums)
	ans := make([]int, l)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j {
				continue
			}
			if nums[j] < nums[i] {
				ans[i]++
			}
		}
	}
	return ans
}
