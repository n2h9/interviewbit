package solution

func diffPossible(A []int, B int) int {
	n := len(A)
	if n <= 1 {
		return 0
	}

	j := 0
	i := 1
	for i < n && j < n {
		for ; i < n && j < i && A[i]-A[j] < B; i++ {
		}
		if i == n {
			break
		}
		for ; j < i && A[i]-A[j] > B; j++ {
		}
		if j == i {
			i++
			continue
		}
		if A[i]-A[j] == B {
			return 1
		}
	}

	return 0
}
