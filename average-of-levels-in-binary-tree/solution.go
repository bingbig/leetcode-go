package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum []int
var count []int

func averageOfLevels(root *TreeNode) []float64 {
	sum = []int{}
	count = []int{}

	traverse(root, 0)
	avg := make([]float64, len(sum))
	for i, _ := range avg {
		avg[i] = float64(sum[i]) / float64(count[i])
	}

	return avg
}

func traverse(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(sum) > level {
		sum[level] += root.Val
		count[level]++
	} else {
		sum = append(sum, root.Val)
		count = append(count, 1)
	}

	level++
	traverse(root.Left, level)
	traverse(root.Right, level)
}
