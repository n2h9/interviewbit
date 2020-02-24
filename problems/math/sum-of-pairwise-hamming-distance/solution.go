package solution

func hammingDistance(A []int) int {
	n := len(A)
	if n <= 1 {
		return 0
	}
	modulo := 1000000007
	var sum, hv int
	var i uint
	for ; i < 32; i++ {
		unitBitsNum := 0
		for j := 0; j < n; j++ {
			if A[j]&(1<<i) > 0 {
				unitBitsNum++
			}
		}
		hv = unitBitsNum * (n - unitBitsNum)
		sum = sum + hv + hv
		for sum > modulo {
			sum -= modulo
		}
	}
	return sum
}
