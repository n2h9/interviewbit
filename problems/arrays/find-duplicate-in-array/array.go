package array

func repeatedNumber(A []int) int {
	n := len(A)
	if n == 0 {
		return -1
	}
	if n == 1 {
		return 1
	}
	b := make([]bool, n-1)
	for i := 0; i < n; i++ {
		if b[A[i]-1] {
			return A[i]
		}
		b[A[i]-1] = true
	}

	return -1
}
