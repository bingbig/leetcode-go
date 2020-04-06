package main

func sum(res *int, n int) bool {
	*res += n
	return n != 0 && sum(res, n-1)
}

func sumNums(n int) int {
	res := 0
	sum(&res, n)
	return res
}
