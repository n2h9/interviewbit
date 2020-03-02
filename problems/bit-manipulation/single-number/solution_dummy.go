package solution

func singleNumberDummy(A []int) int {
	n := len(A)
	s := make(map[int]byte)
	for i := 0; i < n; i++ {
		s[A[i]]++
	}
	for i, v := range s {
		if v == 1 {
			return i
		}
	}

	return 0
}
