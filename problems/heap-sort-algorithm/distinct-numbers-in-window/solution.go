package solution

func dNums(A []int, B int) []int {
	n := len(A)
	if n <= 0 || B > n {
		return []int{}
	}

	res := make([]int, n-B+1)
	i, j := 0, 0
	m := make(map[int]int)

	for ; j < B; j++ {
		m[A[j]]++
	}

	res[i] = len(m)
	i++

	for ; j < n; j++ {
		if m[A[j-B]] == 1 {
			delete(m, A[j-B])
		} else {
			m[A[j-B]]--
		}
		m[A[j]]++
		res[i] = len(m)
		i++
	}

	return res
}
