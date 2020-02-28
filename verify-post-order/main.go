package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	length := len(postorder)
	if length == 0 || length == 1 {
		return true
	}

	root := postorder[length-1]
	i := 0
	for i = 0; i < length-2; i++ {
		if postorder[i] > root {
			break
		}
	}

	// verify right
	for j := i + 1; j < length-1; j++ {
		if postorder[j] < root {
			return false
		}
	}

	if i == 0 {
		return verifyPostorder(postorder[i : length-1])
	}
	if i == length-2 {
		return verifyPostorder(postorder[0 : length-1])
	}

	return verifyPostorder(postorder[0:i-1]) && verifyPostorder(postorder[i:length-1])
}

func main() {
	fmt.Println(verifyPostorder([]int{1, 6, 3, 2, 5}))
	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5}))
	fmt.Println(verifyPostorder([]int{3, 10, 6, 9, 2}))
}
