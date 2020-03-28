package main

import (
	"fmt"
	"math"
)

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	for i := int(math.Sqrt(float64(2 * target))); i >= 2; i-- {
		judge := target - i*(i-1)/2
		if judge%i == 0 {
			begin := judge / i
			temp := make([]int, 0)
			for j := 0; j < i; j++ {
				temp = append(temp, begin+j)
			}
			res = append(res, temp)
		}
	}
	return res
}

func findContinuousSequence2(target int) [][]int {
	seqs := [][]int{}
	s := 1
	max := target/2 + 2
	slice := make([]int, max)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}

	for s < max {
		for l := max - s; l > 1; l-- {
			sum := (s + s + l - 1) * l
			if sum < target*2 {
				break
			}

			if sum == target*2 && s+l <= max {
				seqs = append(seqs, slice[s:s+l])
				break
			}
		}
		s++
	}

	return seqs
}

func main() {
	fmt.Println(findContinuousSequence(9))
	fmt.Println(findContinuousSequence(15))
	fmt.Println(findContinuousSequence(98160))
}
