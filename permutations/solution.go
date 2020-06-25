package solution

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}

	if len(nums) == 2 {
		return [][]int{[]int{nums[0], nums[1]}, []int{nums[1], nums[0]}}
	}

	rest := [][]int{}
	for i := 0; i < len(nums); i++ {
		ps := permute(append(nums[:i], nums[i+1:]...))
		fmt.Println(i, nums[:i], nums[i+1:])
		if ps != nil {
			for _, p := range ps {
				rest = append(rest, append([]int{nums[i]}, p...))
			}
		}
	}

	return rest
}
