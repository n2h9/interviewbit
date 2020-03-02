package solution

func singleNumber(A []int) int {
	n := len(A)
	xor := 0
	for i := 0; i < n; i++ {
		xor ^= A[i]
	}
	return xor
}
