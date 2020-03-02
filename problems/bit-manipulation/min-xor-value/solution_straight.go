package solution

func findMinXorStraight(A []int) int {
	n := len(A)
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return A[0]
	}
	m := 1000000000
	tmp := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp = A[i] ^ A[j]
			if tmp < m {
				m = tmp
			}
		}
	}

	return m
}
