package subarray

func maxSubArray(a []int) int {
	l := len(a)
	sum := make([]int, l)
	sum[0] = a[0]
	for i := 1; i < l; i++ {
		sum[i] = maxI(sum[i-1]+a[i], a[i])
	}
	return maxA(sum)
}

func maxI(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func maxA(a []int) int {
	l := len(a)
	m := 0
	for i := 1; i < l; i++ {
		if a[i] > a[m] {
			m = i
		}
	}

	return a[m]
}
